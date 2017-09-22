package mongo

// 表logs_original_bill
type LogsOriginalBill struct {
	Id              int
	Crawl_id        string
	Platform        string
	Url             string
	Action          string
	Request_header  string
	Request_body    string
	Response_header string
	Response_body   string
	Create_time     string
}
