// Code generated by goctl. DO NOT EDIT.
package types

type UpdateRequest struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	JobCategory string `json:"job_category"`
	JobTypeName string `json:"job_type_name"`
	JobDetail   string `json:"job_detail"`
	JobLocation string `json:"job_location"`
}

type UpdateResponse struct {
}

type RemoveRequest struct {
	Id int64 `json:"id"`
}

type RemoveResponse struct {
}

type CreateRequest struct {
	Cid         int64  `json:"cid"`
	Title       string `json:"title"`
	JobCategory string `json:"job_category"`
	JobTypeName string `json:"job_type_name"`
	JobDetail   string `json:"job_detail"`
	JobLocation string `json:"job_location"`
}

type CreateResponse struct {
	Id int64 `json:"id"`
}

type RefreshRequest struct {
	Cid int64 `json:"cid"`
}

type RefreshResponse struct {
	Id      int64  `json:"id"`
	Company string `json:"company"`
	Number  int64  `json:"number"`
	Brief   string `json:"brief"`
	Picture string `json:"picture"`
}
