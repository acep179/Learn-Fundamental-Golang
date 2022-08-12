package handlers

import (
	dto "dumbmerch/dto/result"
	usersdto "dumbmerch/dto/users"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	// Import validator here ...
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}

	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)}
	json.NewEncoder(w).Encode(response)
}

// Create CreateUser method here ...
// Write this code

//todo Jangan lupa untuk mengatur validasi-nya pada user_request_dto bahwa beberapa field itu required atau harus diisikan

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//ctt dari si struck maupun yang lainnya pada golang itu terdapat kata kunci new yg mana kata kunci new itu seperti kita mereplika atau meng-copy dari struck, pointer-nya dsb. Selayaknya di js juga ada kata kunci new di mana kita menggunakan constructor-nya seperti adanya new Date, new Error, dsb.

	request := new(usersdto.CreateUserRequest)

	//ctt arti dari kodingan di atas adalah saat ini, variabel request itu sudah sama bentukannya dengan struck pada usersdto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		//ctt kodingan di atas adalah utk mengambil value-nya sekaligus memvalidasi apakah yang dikirimkan sudah sesuai dengan struck pada usersdto.CreateUserRequest terkait jenis datanya

		w.WriteHeader(http.StatusBadRequest)

		//ctt jika terjadi error (tidak sesuai dengan kerangka pada usersdto.CreateUserRequest) maka status yg dikirimkan adalah Bad Request alias data yg dikirimkan itu tidak sesuai dengan kerangka pada usersdto.CreateUserRequest alias ada tipe data yg tdk sesuai, dll.

		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//ctt kita juga membutuhkan validatornya maka kita buatkan sebuah variabel dengan nama validation yg akan menyimpan value dari validator-nya

	validation := validator.New()

	//ctt setelah dibuatkan variabel-nya, maka skrg kita tinggal validasi struck-nya apakah data yg dikirmkan sesuai

	err := validation.Struct(request)

	//ctt err nya didapatkan dari validation-nya adapun yg divalidasi adalah terkait required-nya

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//todo mengirimkan data request ke repository

	//ctt Karena repository meminta datanya dalam bentuk struck models.User. Maka, kita siapkan terlebih dahulu sebagaimana di bawah ini
	//ctt yaitu, kita convert ke struck models.User

	// data form pattern submit to pattern entity db user
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	//ctt karena kita tidak memiliki data id, created_at, dan updated_at. Maka, kita isikan seperlunya saja. Kita isikan hanya dengan data yang kita miliki saja, yaitu name, email, dan password.

	//ctt Sehingga sekarang kita sudah memiliki data dalam bentuk struck-nya models.User dan tinggal kita kirimkan data tersebut ke repository karena kita akan menggunakan kontrak dari repository-nya karena ia memintanya dalma bentuk models.User.

	//todo di bawah ini adalah cara untuk mengirimkan data inputan dari handler ke repository

	data, err := h.UserRepository.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

//todo Setelah selesai melakukan konfigurasi pada handler-nya. Maka, sekarang kita masukkan handler ini ke routes-nya.

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
