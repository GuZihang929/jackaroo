// Code generated by goctl. DO NOT EDIT.
package types

type JobsRequest struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	JobCategory string `json:"job_category"`
	JobTypeName string `json:"job_type_name"`
	JobLocation string `json:"job_location"`

	NextId int64	`json:"next_id"`
	PageSize int64	`json:"page_size"`
}

type JobsResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Date []Job  `json:"date"`
}

type TagAddRequest struct {
	Tag string `json:"tag"`
}

type TagAddResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Date string `json:"date"`
}

type CompanySumResponse struct {
	Code int64      `json:"code"`
	Msg  string     `json:"msg"`
	Date []Companys `json:"date"`
}

type JobRequest struct {
	Id int64 `json:"id"`
}

type JobResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Date Job    `json:"date"`
}

type Job struct {
	Id          int64  `json:"id"`
	CId         string  `json:"c_id"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	JobCategory string `json:"job_category"`
	JobTypeName string `json:"job_typeName"`
	JobDetail   string `json:"job_detail"`
	JobLocation string `json:"job_location"`
	Push_time   string `json:"push_time"`
	Fetch_time  string `json:"fetch_time"`
}

type Companys struct {
	Id	string	`json:"id"`
	Company string `json:"company"`
	Number     int64 `json:"number"`
	Brief	string	`json:"brief"`
	Picture	string	`json:"picture"`
}
