package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

type errObj struct {
	Error string `json:"error"`
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		myErrObj := errObj{
			err.Error(),
		}
		JSON(w, statusCode, myErrObj)
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
