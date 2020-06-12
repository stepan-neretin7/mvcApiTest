package services

import "net/http"

func BadRequest(w http.ResponseWriter, err error){
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}


func CheckRequestMethod(req http.Request, need string) bool {
	if req.Method != need {
		return false
	}
	return true
}