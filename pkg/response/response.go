package appRes

import "victoria-falls/pkg/response/pagination"

// MetaResponse types
type MetaResponse struct {
	Error   uint   `json:"error"`
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// APIResponse types
type APIResponse struct {
	Meta  *MetaResponse `json:"meta,omitempty"`
	Data  interface{}   `json:"data"`
	Error interface{}   `json:"error"`
}

// Detail ...
type Detail struct {
	Detail interface{} `json:"detail"`
}

// List  ...
type List struct {
	List       interface{}          `json:"list"`
	Pagination *pagination.Response `json:"pagination,omitempty"`
}

// DetailList ...
type DetailList struct {
	Detail     interface{}          `json:"detail"`
	List       interface{}          `json:"list"`
	Pagination *pagination.Response `json:"pagination,omitempty"`
}

type Response struct {
	Name       string
	Status     bool
	Messages   string
	Code       int
	ErrorCode  uint
	Errors     interface{}
	Result     interface{}
	Detail     interface{}
	List       interface{}
	Meta       interface{}
	Pagination *pagination.Response
}

// Result ...
func Result(res *Response) (int, APIResponse) {
	var data interface{}
	if res.Detail != nil && res.List != nil {
		data = DetailList{res.Detail, res.List, res.Pagination}
	} else if res.Detail != nil {
		data = Detail{
			res.Detail,
		}
	} else if res.List != nil {
		data = List{
			res.List,
			res.Pagination,
		}
	} else {
		data = res.Result
	}
	response := APIResponse{
		Meta: &MetaResponse{
			Error:   res.ErrorCode,
			Code:    res.Code,
			Status:  res.Status,
			Message: res.Messages,
		},
		Data:  data,
		Error: res.Errors,
	}
	return res.Code, response
}
