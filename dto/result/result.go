package dto

// Declare SuccessResult struct here ...
type SuccessResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Declare ErrorResult struct here ...
type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
