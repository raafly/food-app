package exception

import (
	"net/http"

	helper "github.com/raafly/food-app/pkg/helpers"
)

type ErrResponse struct {
	Code	int
	Status	string
	Message	string
}

func ErrorHadler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok:= err.(NotFoundError)	
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := ErrResponse {
			Code: http.StatusNotFound,
			Status: "NOT FOUND",
			Message: exception.Error,
		}

		helper.WriteToRequestBody(w, webResponse)
		
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request ,err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := ErrResponse {
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Message: "SERVER IS BUSY",
	}

	helper.WriteToRequestBody(w, webResponse)
}