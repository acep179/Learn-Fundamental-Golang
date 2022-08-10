package usersdto

//ctt Setiap data yg ditransfer baik dlm proses request maupun response, akan melalui dto ini. Yaitu, Data Transfer Object.
//ctt Karena kita akan melakukan fetching, maka kita membutuhkan response dari bentukan sruct user-nya

// Declare UserResponse struct here ...
type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

//ctt Karena kita akan mem-fetching data, maka format data yg kita ambil hanya ID, Name, Email, dan Password.
