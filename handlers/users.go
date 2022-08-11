package handlers

import (
	dto "dumbmerch/dto/result"
	usersdto "dumbmerch/dto/users"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//todo Declare handler struct here ...
//ctt Langkah pertama yg kita lakukan adalah membuat handler struct-nya yaitu dengan mengambilnya dari package repositories => function UserRepository yg berupa interface. Pastikan bahwa func tsb merupakan func global.
type handler struct {
	UserRepository repositories.UserRepository
	//ctt Maksud dari kodingan di atas adaah: utk ke depannya, kita cukup memanggilnya dengan UserRepository saja
}

//todo Declare HandlerUser function here ...
func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
	//ctt utk si handler ini, nantinya akan kita teruskan ke si repositories-nya. Otomatis, kita isikan utk si Repositories-nya ini berasal dari UserRepository
	//ctt adapun func HandlerUser itu akan kita gunakan ketika kita membuat Routes-nya
}

//todo Declare FindUsers method here ...
func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
	//ctt jadi, ini adalah sebuah Method handler yg sudah berisikan value terkait repository mana yg akan meng-handle proses selanjutnya (handler di atas)

	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()
	//ctt karena di sini dia kita jadikan Method, maka langsung kita tuliskan prosesnya di atas. Maka, prosesnya kita tuliskan yaitu berasal dari UserRepository-nya yg mana Methodnya akan disesuaikan dengan yg ada di Repository-nya yaitu .FindUsers()
	//ctt adapun nilai returnya itu ada dua yaitu nilainya dan error-nya sesuai dgn apa yg kita tuliskan di repository-nya.

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		//ctt Di sini, kita sesuaikan utk responnya agar sesuai dengan respon yang kitainginkan
		//ctt Adapun respon yang diberikan tidak hanya respon success saja, akan tetapi juga termasuk respon error-nya. Maka dari itu, kita perlu menyiapkan kerangka kedua respon tersebut pada DTO agar baik ketika error maupun success maka responnya menjadi sama semua.
		//ctt Kita siapkan kerangkanya pada folder dto/result/result.go

		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

// Declare GetUser method here ...
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

// Declare convertResponse function here ...
func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
