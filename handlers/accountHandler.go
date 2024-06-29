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

func GetAccountHandler(responseW http.ResponseWriter, request *http.Request) {
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

func RegisterNewUserHandler(responseW http.ResponseWriter, request *http.Request) {
	// get over body request
	var newUser model.User
	err := json.NewDecoder(request.Body).Decode(&newUser)

	if err != nil {
		http.Error(responseW, "Invalid payload "+err.Error(), http.StatusBadRequest)
		return
	}
	// check if username exists
	res, err := util.CheckUsernameExist(newUser.Username)

	if err != nil {
		http.Error(responseW, "Invalid payload - username check "+err.Error(), http.StatusBadRequest)
		return
	}

	if res {
		http.Error(responseW, "Invalid payload - username exists ", http.StatusBadRequest)
		return
	}

	rawPassword := newUser.RawPassword
	hashedPassword, err := util.HashPassword(rawPassword)

	if err != nil {
		http.Error(responseW, "Invalid payload - password "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.InsertNewAccount(newUser, hashedPassword)

	if err != nil {
		http.Error(responseW, "Invalid payload - db transaction "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(responseW, "New user id = ", id)
}

func AccountLoginHandler(responseW http.ResponseWriter, request *http.Request) {
	var user model.User
	json.NewDecoder(request.Body).Decode(&user)
	username, raw_password := user.Username, user.RawPassword

	// check if username exist
	check, err := util.CheckUsernameExist(username)
	if err != nil {
		http.Error(responseW, "Invalid payload - username check "+err.Error(), http.StatusBadRequest)
		return
	}
	if !check {
		http.Error(responseW, "Invalid payload - username does not exist ", http.StatusBadRequest)
		return
	}

	// check password
	check, err = util.CheckRawPassword(raw_password, username)
	if err != nil {
		http.Error(responseW, "Invalid payload - username password check "+err.Error(), http.StatusBadRequest)
		return
	}
	if !check {
		http.Error(responseW, "Invalid payload - WRONG password ", http.StatusBadRequest)
		return
	}

	tokenStr, err := util.CreateJWTToken(username)
	if err != nil {
		http.Error(responseW, "Error token not successful --- "+err.Error(), http.StatusBadRequest)
		return
	}
	responseW.Header().Set("Content-Type", "application/json")
	responseW.WriteHeader(http.StatusOK)
	fmt.Fprint(responseW, tokenStr)
}
