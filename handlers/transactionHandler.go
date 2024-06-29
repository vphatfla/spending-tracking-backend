package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	customerror "github.com/spending-tracking/customError"
	"github.com/spending-tracking/db"
	"github.com/spending-tracking/model"
	"github.com/spending-tracking/util"
)

func GetAllTransactionByUserIdHandler(responseW http.ResponseWriter, request *http.Request) {
	responseW.Header().Set("Content-Type", "application/json")
	tkCheck, err := util.TokenRequestHandling(request)

	if err != nil {
		if errors.Is(err, customerror.NoAuthError()) {
			responseW.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(responseW, "No Authorization detected")
			return
		}
		if errors.Is(err, customerror.InvalidJWTToken()) {
			responseW.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(responseW, "Invalid token")
			return
		}
	}
	if !tkCheck {
		responseW.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(responseW, "Unable to authorize for resources")
		return
	}
	// params
	query := request.URL.Query()
	userIdStr := query.Get("user_id")

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
	responseW.Header().Set("Content-Type", "application/json")
	tkCheck, err := util.TokenRequestHandling(request)

	if err != nil {
		if errors.Is(err, customerror.NoAuthError()) {
			responseW.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(responseW, "No Authorization detected")
			return
		}
		if errors.Is(err, customerror.InvalidJWTToken()) {
			responseW.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(responseW, "Invalid token")
			return
		}
	}
	if !tkCheck {
		responseW.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(responseW, "Unable to authorize for resources")
		return
	}
	// params
	var transaction model.Transaction
	err = json.NewDecoder(request.Body).Decode(&transaction)
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
