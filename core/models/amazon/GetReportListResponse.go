package amazon

import "time"

type GetReportListResponse struct {
	GetReportListResult struct {
		HasNext string `json:"HasNext"`
		ReportInfo []struct {
			ReportType      string    `json:"ReportType"`
			Acknowledged    string    `json:"Acknowledged"`
			ReportID        string    `json:"ReportId"`
			ReportRequestID string    `json:"ReportRequestId"`
			AvailableDate   time.Time `json:"AvailableDate"`
		} `json:"ReportInfo"`
	} `json:"GetReportListResult"`
	ResponseMetadata struct {
		RequestID string `json:"RequestId"`
	} `json:"ResponseMetadata"`
}
