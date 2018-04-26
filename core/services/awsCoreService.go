package services

import (
	"fmt"
	"gopkg.in/xmlpath.v2"
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)
const (
	apiVersion                 = "1"
	LIST_ORDER                 = "ListOrders"
	REQUEST_REPORT             = "RequestReport"
	GET_REPORT_LIST            = "GetReportList"
	GET_REPORT                 = "GetReport"
	POST                       = "POST"
	GET                        = "GET"
	CREATE_INVOICE_SOCKET_FLOW = "https://sokt.io/FH4S7ryuTn7u5SLwbVgu/amazon-flow-amazon-create-invoice"
)
var client Client
func CallListReportApi (reportId string, creds AwsCreds) (string){
	duration := time.Minute
	time.Sleep(duration)
	var repoId string
    for {
	    AwsClientForReportList := NewClient(client)
	    AwsClientForReportList.Operation = GET_REPORT_LIST
	    AwsClientForReportList.Action = GET_REPORT_LIST
	    AwsClientForReportList.Method = POST
	    AwsClientForReportList.AwsCreds = creds
	    req, err := AwsClientForReportList.RequestForReportList(reportId)
	    if (err != nil) {
		    fmt.Println(req)
		    fmt.Println(err)
	    }
	    awsPostClient := &http.Client{}
	    response, err := awsPostClient.Do(req)
	    if err != nil {
		    fmt.Println(err)
	    }
	    bodyOfReportList, err := ioutil.ReadAll(response.Body)
	    if err != nil {
		    fmt.Println(bodyOfReportList)
		    fmt.Println(err)
	    }
	    newPath := xmlpath.MustCompile("/GetReportListResponse/GetReportListResult/ReportInfo/ReportId")
	    root, err := xmlpath.Parse(bytes.NewReader(bodyOfReportList))
	    if err != nil {
		    fmt.Println("Error here -2")
	    }
	    //fmt.Println(root)

	    if value, ok := newPath.String(root); ok {
		    repoId = value
	    }

	    fmt.Println("my",repoId)
	    if(len(repoId)!=0){break}
    }
  return repoId
}
