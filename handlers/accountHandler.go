package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func AccountHandler(responseW http.ResponseWriter, request *http.Request) {
	// params
	query := request.URL.Query()

	userIdStr := query.Get("id")

	//responseW.Header().Set("Content-Type", "application/json")

	// empty or invalid id
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		responseW.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(responseW, "invalid userId")
		return
	}

	fmt.Fprint(responseW, "userId = ", userId, " \n")
}
