package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spending-tracking/db"
	"github.com/spending-tracking/model"
)

func GetAllTransactionByUserIdHandler(responseW http.ResponseWriter, request *http.Request) {
	// params
	query := request.URL.Query()
	userIdStr := query.Get("user_id")
	responseW.Header().Set("Content-Type", "application/json")

	// empty or invalid id
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		responseW.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(responseW, "invalid userId")
		return
	}

	transactions, err := db.GetAllTransactionByUserId(userId)
	if err != nil {
		responseW.WriteHeader(http.StatusBadRequest)
		http.Error(responseW, err.Error(), http.StatusBadRequest)
		return
	}

	transactionJson, err := json.Marshal(transactions)
	if err != nil {
		responseW.WriteHeader(http.StatusBadRequest)
		http.Error(responseW, err.Error(), http.StatusBadRequest)
		return
	}

	responseW.WriteHeader(http.StatusOK)
	fmt.Fprint(responseW, string(transactionJson))
}

func PostNewTransactionHandler(responseW http.ResponseWriter, request *http.Request) {
	// params
	var transaction model.Transaction
	err := json.NewDecoder(request.Body).Decode(&transaction)
	if err != nil {
		http.Error(responseW, "Invalid payload   "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.UploadTransaction(transaction)

	if err != nil {
		http.Error(responseW, err.Error(), http.StatusBadRequest)
		return
	}
	// empty or invalid id
	responseW.Header().Set("Content-Type", "application/json")
	fmt.Fprint(responseW, "new transaction id = ", id)
}
