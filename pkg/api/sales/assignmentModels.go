package sales

import common2 "github.com/erply/api-go-wrapper/pkg/api/common"

//saveAssignmentResponse ...
type saveAssignment struct {
	Status  common2.Status `json:"status"`
	Records []struct {
		AssignmentID int64 `json:"assignmentID"`
	} `json:"records"`
}
