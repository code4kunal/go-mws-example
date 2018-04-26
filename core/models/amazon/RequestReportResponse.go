package amazon

import (

)


type RequestReportResponse struct {
		Xmlns               string `xml:"-xmlns,attr"`
		RequestReportResult struct {
			ReportRequestInfo struct {
				ReportType             string    `xml:"ReportType"`
				ReportProcessingStatus string    `xml:"ReportProcessingStatus"`
				//EndDate                time.Time `xml:"EndDate"`
				Scheduled              string    `xml:"Scheduled"`
				ReportRequestID        string    `xml:"ReportRequestId"`
				//SubmittedDate          time.Time `xml:"SubmittedDate"`
				//StartDate              time.Time `xml:"StartDate"`
			} `xml:"ReportRequestInfo"`
		} `xml:"RequestReportResult"`
		ResponseMetadata struct {
			RequestID string `xml:"RequestId"`
		} `xml:"ResponseMetadata"`
}
