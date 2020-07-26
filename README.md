ERPLY API Go SDK
--------
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/erply/api-go-wrapper/pkg/api?tab=doc)
[![API Reference](https://img.shields.io/badge/api-reference-blue.svg)](https://learn-api.erply.com/)

This SDK covers the [ERPLY API](https://erply.com/erply-api/) requests. 

Client Structure
------
Majority of the request wrappers are available through the client.
The client is described in `GoDoc` type `Client` and in `/pkg/api/client.go`. It is divided into sub-clients for each topic that the underlying API covers. 
For now not all the requests are mapped to topics. Such request wrappers are in `/pkg/api` directory. 
Some of the requests are accessible not from the client, but from the `auth` package of this SDK. They are covered in the example in `/examples` directory.

Install
-------
   `go get github.com/erply/api-go-wrapper@X.Y.Z`
   
   where X.Y.Z is your desired version.
   
Clients
--------
Ways of using the API: 
* One is you create a `Partner Client` that will always use the partner key with requests and will have access to the requests that require the partner key.
* You can use the simple `Client` that will work without the partner key also.
* You can also create a client that can act like a partner client, normal one and it is possible to define the headers that will be added for every request on your own. For that one please use the `NewClientWithCustomHeaders` constructor

You can find the example in the `/examples` directory for the client initialization process

Contributing
-------
This library is not in the final state and it means for continuous development. Therefore I would like to cover the entire ERPLY API and contributions are of course always welcome. The calling pattern is pretty well established, so adding new methods is relatively straightforward. 

Advanced listing
--------
### Overview
Advanced listing was disigned to read large data collections by multiple parallel fetchers with respect of API limitations. Moreover this API will use the minimal amount of requests by packing them into bigger bulk API calls, so the too many request failures will be less probable. 

Generally all you need is to create a `Lister` struct giving `ListingSettings`which customises the parallel processing. 
Then you can call either `Get` or `GetGrouped` method, they will give you in return a channel of items which you can consume concurrently with go routines. The fetchers of the library will close the channel once all required data was fetched from a corresponding API, so you can securely iterate over the channel with your go routines.

 
### Getting started
In this intro we want to read all products (`products.Product`) which were changed after a defined date. Since we expect a large amount of such items, we would prefer to use a parallel listing API.

**Let's create a `Lister` struct:**

    productsDataProvider := products.NewListingDataProvider(cl.ProductManager) #we define here that we call lister against Products API
    
    lister := sharedCommon.NewLister(
        sharedCommon.ListingSettings{
            MaxRequestsCountPerSecond: 5, # respecting our expected limitation of 300 request per minute
            StreamBufferLength:        10, # it's a buffer of the output product channel
            MaxItemsPerRequest:        300, #at a time each fetcher will make one request with 3 bulk subrequests per 100 items, this number cannot be more than 10000 (max 100 bulk requests x with max 100 items per request)
            MaxFetchersCount:          10, # the amount of parallel fetchers to read data from API
        },
        productsDataProvider,
        func(sleepTime time.Duration) {
            time.Sleep(sleepTime) # this is needed to customize sleep logic
        },
    )

**Let's define the processing context with a timeout**  

 
        ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60) #timeout if the whole process takes > 60s
        defer cancel()


**Let's start fetching and receive our results channel**
        
        prodsChan := lister.Get(ctx, map[string]interface{}{
            "changedSince": time.Date(2020, 2, 15, 0, 0, 0, 0, time.UTC).Unix(), # filter the products changed after 2020-02-15
        })

**Let's create our consumer for the product channel**

    prods := make([]products.Product, 0) #the struct where we save all products
    
    for prod := range prodsChan {
        if prod.Err != nil { # API errors are appearing here
            panic(prod.Err)
        }
        prods = append(prods, prod.Payload.(products.Product))
    }
    
    fmt.Printf("%+v", prods) #we nicely print here the output but of course in a real application you would do some processing of products
   
That is pretty it. The full example you can find in the examples folder (see `examples/products/main.go` `GetProductsInParallel` method)

### Get vs GetGrouped API methods
You can call `Get` or `GetGrouped` methods on the `Lister` struct:

     prodsChan := lister.Get(ctx, map[string]interface{}{
         "changedSince": time.Date(2020, 2, 15, 0, 0, 0, 0, time.UTC).Unix(), # filter the products changed after 2020-02-15
     })
     #vs
     prodsGroupedChan := lister.GetGrouped(ctx, map[string]interface{}{
         "changedSince": time.Date(2020, 2, 15, 0, 0, 0, 0, time.UTC).Unix(), # filter the products changed after 2020-02-15
     }, 10)
    
In the first case you will get a flat channel of products like
 
    {product1, product2, ...productN}

in the second case the result will be 

    {[]{product1, ...productX}, []{productX+1...productY}, []{productY+1...productZ}, ...[]GroupGroupN}
    
Why the `GetGrouped` method is needed?

Imagine a situation, where for each item you get from the Listing API you would need to make a subrequest to fetch additional details, e.g. in case with the products API for each product and corresponding `SupplierID` you would like to have a supplier name. 

An obvious solution would be to fetch products channel from the `Get` method, start consumption and for each `Product` take `SupplierID` value and call the [getSuppliers](https://learn-api.erply.com/requests/getsuppliers) API. Unfortunately this will require too many individual calls and will probably hit the api requests limit. 

Instead of this you could call the `GetGrouped` method to pack individual products into groups of e.g. 100 items, so on each iteration you can make a bulk request with 100 subrequests (since you can filter only one supplier by id per subrequest). Let's demonstrate it:

     prodsGroupedChan := lister.GetGrouped(ctx, map[string]interface{}{
         "changedSince": time.Date(2020, 2, 15, 0, 0, 0, 0, time.UTC).Unix(), # filter the products changed after 2020-02-15
     }, 100) # set groupSize to 500 to get from channel 500 grouped products per iteration 
     
     supplierCli := cl.CustomerManager #supplier API
     	
     supplierBulkFilter := make([]map[string]interface{}, 0, 100) #here we a bulk filter with 100 subrequests
     supplierNames := make([]string, 0) #here is our final result with all supplier names
     
     for prodGroup := range prodsGroupedChan { #starting consuming the channel with groups of products
        for _, prod := range prodGroup { #each prodGroup contains max 100 products
            if prod.Err != nil {
                panic(prod.Err) #handling API errors
            }
            supplierBulkFilter = append(supplierBulkFilter, map[string]interface{}{
                "supplierID": prod.Payload.(products.Product).SupplierID, //todo check for payload type
            })
        }
 
        supplierRespBulk, err := supplierCli.GetSuppliersBulk(ctx, supplierBulkFilter, map[string]string{}) #fetching 100 suppliers in one bulk request
        if err != nil {
            panic(err)
        }
        //todo check supplierRespBulk for status
        
        for _, supplierItem := range supplierRespBulk.BulkItems {
            supplierNames = append(supplierNames, supplierItem.Suppliers[0].FullName) #pack the result into the output slice
        }
     }

### Shareable requests throttler
     
You might ask yourself what happens if the product lister makes max 5 requests per second, and the consumer, as it's running in parallel, will make the 6th request which will lead to a "too many requests" failure. That's absolutely true. If your application code is running in parallel with the lister, you should also take into account the requests speed of it.

To limit the requests count, the `Lister` uses the `Throttler` interface. The library offers `SleepThrottler` as it's implementation. It's logic is pretty simple: each fetcher instance will call `Throttle` method of the `SleepThrottler` before any call to Erply API. `SleepThrottler` will count the amount of requests in the last second, and if it's more than the defined count (e.g. 5), it will sleep 1 second. As a result, too fast requests will be slowed down if needed. 

To solve our problem we need to make sure that your application code is using the same throttler, so it should be practically shared with the lister. To achieve this we can do the following:

    sleepFunc := func(sleepTime time.Duration) {
       time.Sleep(sleepTime)
    }
    thrl := NewSleepThrottler(settings.MaxRequestsCountPerSecond, sleepFunc)
    
    lister := sharedCommon.NewLister(
        sharedCommon.ListingSettings{
            MaxRequestsCountPerSecond: 5,
            StreamBufferLength:        10,
            MaxItemsPerRequest:        300,
            MaxFetchersCount:          10,
        },
        productsDataProvider,
        sleepFunc,
    )
    lister.SetRequestThrottler(thrl) #this is a new code, we give our shareable throttler to the lister
    
    prodsGroupedChan := lister.GetGrouped(ctx, map[string]interface{}{}, 100) 
         
    supplierCli := cl.CustomerManager
          	
    supplierBulkFilter := make([]map[string]interface{}, 0, 100)
    supplierNames := make([]string, 0)
      
    for prodGroup := range prodsGroupedChan { #starting consuming the channel with groups of products
        for _, prod := range prodGroup { #each prodGroup contains max 100 products
            if prod.Err != nil {
                panic(prod.Err) #handling API errors
            }
            supplierBulkFilter = append(supplierBulkFilter, map[string]interface{}{
                "supplierID": prod.Payload.(products.Product).SupplierID, //todo check for payload type
            })
        }
        
        thrl.Throttle() #this is a new code, here we make sure that we don't hit the request limit with the lister
        supplierRespBulk, err := supplierCli.GetSuppliersBulk(ctx, supplierBulkFilter, map[string]string{})
        if err != nil {
            panic(err)
        }
    
         for _, supplierItem := range supplierRespBulk.BulkItems {
             supplierNames = append(supplierNames, supplierItem.Suppliers[0].FullName) #pack the result into the output slice
         }
  }
  
### Configuration hints
As you already might have noticed, the main configuration data is passed in the `ListingSettings` struct:

        sharedCommon.ListingSettings{
            MaxRequestsCountPerSecond: 5,
            StreamBufferLength:        10,
            MaxItemsPerRequest:        300,
            MaxFetchersCount:          10,
        }


**MaxRequestsCountPerSecond**

This option indicates the amount of requests per second, which are allowed for the `Lister`. Practically the `SleepThrottler` uses this number to identify too often requests and to trigger sleeping logic, if your fetchers are hitting the defined limit. 

**StreamBufferLength**
 This indicates the buffer length of the channel, which is returned from the `Get` or `GetGrouped` methods of the `Lister`. To select the correct value for this parameter, you need to consider the fetchers count you set in the `MaxFetchersCount` parameter, the amount of the output consumers and the difference in publishing and consumption speed. 
 
 Imagine that you fetch products with 10 fetchers and each is sending 10 products per second giving in total 100 prod/s speed. At the same time your consumer(s) process 500 prod/s, so you probably don't need to set `StreamBufferLength` to a very high number. On the other hand if your consumer is able to process 1 product per second, it will block all 10 fetchers, so you won't have a good processing speed. In this case the buffer length should be a number around 100, so you will let your consumer to do most of the work after the fetchers finish their work. 
 
**MaxItemsPerRequest**
This parameter sets the limit of total items, which are allowed to be fetched per bulk request. One might ask, why this parameter is needed, since we could always use the max possible fetching count per request: 100 bulk requests x 100 per request. Unfortunately we noticed that fetching of 10 000 items per bulk request could take up to 30 seconds, but making 20 requests with 500 items each can take up to 10 seconds.

You should set a balanced value for this option depending on your filters, API type and overall load of the Erply infastructure, try to set this value between 500 and 1000 and see which one gives the most performance. 

**MaxFetchersCount**
The amount of parallel fetchers, which are allowed to call the Erply API. Don't try to use very high numbers here, hoping to reach the most of the processing speed. This value should be well balanced with the consumption speed and `StreamBufferLength` parameter. 

For example if you fetchers are able to send 10 items per second each but your consumers can process only 1 item per second, you probably would need only one consumer and `StreamBufferLength` = 10. On the other hand if you can consume 100 items per second, set the `MaxFetchersCount` = 10 and `StreamBufferLength` = 10.

### Implementation details

The `Lister` is based on a popular [Fan-out](https://blog.golang.org/pipelines) concurrent pattern where multiple go routines are getting payload from a single input channel. An output channel is created for each go routine, where it sends the result of the parallel work. There is also a separate go routine which consumes from all those channels and sends the result to the single output channel, which is returned in the output of the `Get` or `GetGrouped` method.

### Get method algorithm

`Lister` requires `DataProvider` interface which wraps all individual Erply APIs which should support advanced parallel listing. A corresponding API should be able to give the amount of elements which match the filter (usually we use `RecordsTotal` value of a [bulk status response](https://learn-api.erply.com/getting-started/bulk-api-calls)) and support pagination options `recordsOnPage`, `pageNo` (see e.g. [getProducts](https://learn-api.erply.com/requests/getproducts)).

#### Planning stage
`Lister` calls  `DataProvider` with the provided filters to figure out how many elements are matching the filters. This also helps to stop the execution and close output channel once all items are fetched.

#### Creating Cursor channel and planning the fetching
`Lister` creates an input channel of `[]Cursor` items, where each `Cursor` identifies the chunk of the data which should be fetched from an API. 

The idea of the cursor is simple: imagine you have 100 products in total. You allow to fetch 10 items per request (see `MaxItemsPerRequest` option). You would need 10 cursors to fetch the whole sequence: 

`[1-10], [11-20], [21-30], [31-40], [41-50], [51-60], [61-70], [71-80], [81-90], [91-100]`

You can simplify cursor data as the offset + limit numbers e.g.

`[1, 10], [11, 10], [31, 10] etc.`

Knowing the amount of items per request and the total number of items to process, we can calculate all the Cursors, which are needed to fetch all items respecting the request limit. 

But why we create `chan []Cursor` rather than `chan Cursor`? Of course we could create 10 cursors and make 10 parallel requests but we have bulk request, so practically we could send 1 request with 10 subrequests and spare 9 network requests which will speed up our execution. So the algorithm takes into account the max number of items per 1 API request (100) and the `MaxItemsPerRequest` value to understand how many bulk requests are needed to fetch the whole amount of items.

Let's consider following more realistic example:

total items count: 100 000
max items per request: 500
chan []Cursor: first item send over the channel will look like []Cursor{{offset:0, limit: 100}, {offset:100, limit: 100}, {offset:200, limit: 100}, {offset:300, limit: 100}, {offset:400, limit: 100}}

This means the first element would require to make a bulk request with 5 subrequests respecting the limit of 500. To total amount of items transferred through the `[]Cursor` channel will be 100 000 / 500 = 20. Of course the logic controls that the amount of `[]Cursor` items is not greather than 100 (you cannot make more than 100 subrequests in one bulk request)

### Starting fetching
We start `MaxFetchersCount` go routines (see `fetchItemsChunk` method). Each go routine creates an output channel (`chan Item`) and returns it to the main process. Each go routine starts consumption from the `chan []Cursor` channel.
It creates a bulk request with len([]Cursor) items and applies the filters `map[string]interface{}`, which was passed to the Lister in the `Get` method. 

It calls a bulk request respecting the throttling logic. It sends each item to the output channel (`chan Item`) and reads the next cursor chunk from the input channel. 

All fetchers are respecting context cancellation when consuming the `[]Cursor` channel. Once the cursors are send to the channel, it will be closed, so the fetchers will exit the channel loop and close their output channels. 

If a fetcher will get some unhealthy status from a bulk requests or will have any other non-recoverable failure, it will send an `Item` with a non-empty error field to the output channel. It will also stop further execution. The external application logic should handle the error and practically stop all the execution of the whole app. Currently library doesn't support error recovery and restart/reconnect policy.

### Merging output channels
We create a single output channel, which will at the end be returned to the `Get` method caller. For each of the go routines output channel we start another go routine to consume from it. Each such routine will just forward all consumed items to the single output channel. 

By using the WaitGroup struct, it will start another go routine which will close the final output channel once the all inbound output channels are closed.
The merged output channel is returne to the `Get` method caller. 
