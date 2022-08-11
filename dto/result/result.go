package dto

//todo Declare SuccessResult struct here ...
type SuccessResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

//ctt Jika success maka kita akan memanggil code dan datanya

//todo Declare ErrorResult struct here ...
type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//ctt Jika error maka kita akan memanggil code dan message-nya
//ctt Kedua bentuk respon di atas akan dikirimkan pada bagain JSON-nya.
