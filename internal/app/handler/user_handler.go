package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BoruTamena/habitant_hub/internal/app/model"
	"github.com/BoruTamena/habitant_hub/internal/app/service"
)

type Userhandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *Userhandler {

	return &Userhandler{userService}

}

func (uh *Userhandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // flag as a bad reques
		return
	}

	err = uh.userService.CreateUserService(&user)

	if err != nil {

		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}
