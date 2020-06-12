package api

import (
	"encoding/json"
	"errors"
	"mvcApiTest/pkg/models"
	"mvcApiTest/pkg/repositories"
	"mvcApiTest/pkg/services"
	"net/http"
)

func writeResponse(w http.ResponseWriter, data []byte){
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(data)
	if err != nil{
		panic(err)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request){
	userService := services.NewUserService(repositories.NewUsersRepository())
	resp, err := json.Marshal(userService.GetAll())
	if err != nil {
		services.BadRequest(w, err)
	}
	writeResponse(w, resp)
}

func CreateUsers(w http.ResponseWriter, r *http.Request){
	userService := services.NewUserService(repositories.NewUsersRepository())
	if !services.CheckRequestMethod(*r, "POST"){
		services.BadRequest(w, errors.New("this method for this route is not supported"))
	}
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		services.BadRequest(w, err)
	}
	answer := make(map[string]interface{})
	userCount, err := userService.IsUserExists(user.Email)
	if err != nil{
		services.BadRequest(w, err)
	}
	if userCount > 0{
		answer["message"] = "user with this email already exists"
		data, err := services.GenerateResponse(answer)
		if err != nil{
			services.BadRequest(w, err)
		}
		writeResponse(w, data)
		return
	}
	_ = userService.InsertUser(user)

	answer["message"] = "ok"
	data, err := services.GenerateResponse(answer)
	if err != nil{
		services.BadRequest(w, err)
	}
	writeResponse(w, data)
}