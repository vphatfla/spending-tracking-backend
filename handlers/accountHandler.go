package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spending-tracking/db"
)

func GetAccountHandler(responseW http.ResponseWriter, request *http.Request) {
	// params
	query := request.URL.Query()

	userIdStr := query.Get("id")

	responseW.Header().Set("Content-Type", "application/json")

	// empty or invalid id
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		responseW.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(responseW, "invalid userId")
		return
	}

	user, err := db.GetAccountById(userId)
	if err != nil {
		responseW.WriteHeader(http.StatusBadRequest)
		http.Error(responseW, err.Error(), http.StatusBadRequest)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		responseW.WriteHeader(http.StatusBadRequest)
		http.Error(responseW, err.Error(), http.StatusBadRequest)
		return
	}

	responseW.WriteHeader(http.StatusOK)
	fmt.Fprint(responseW, string(userJson))
}

func InsertNewUserHandler(responseW http.ResponseWriter, request *http.Request) {
	// hashedPassword, err := util.HashPassword(passwordStr)

	// if err != nil {
	// 	return insertedTransaction, err
	// }
}
