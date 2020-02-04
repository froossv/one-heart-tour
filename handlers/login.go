package handlers

import (
	"encoding/json"
	"net/http"
)

//Login User
func Login(w http.ResponseWriter, r *http.Request) {
	var user, dbUser User
	var resp Response
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Error"
		return
	}

	db := GetDB()
	defer db.Close()

	db.Find(&dbUser, "username = ?", user.Username)
	if dbUser.Password == user.Password {
		w.WriteHeader(http.StatusOK)
		resp.Message = "Success"
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		resp.Message = "Error"
	}

	respJSON, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)

}
