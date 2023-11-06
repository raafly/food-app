package auth

import (
	"encoding/json"
	"net/http"

	"github.com/raafly/food-app/internal/listing"
	helper "github.com/raafly/food-app/pkg/helpers"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{} 
	if err := decoder.Decode(&data); err != nil {
		// error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := listing.WebResponse {
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToRequestBody(w, webResponse)
	} else {
		// oke 
		middleware.Handler.ServeHTTP(w, r)
	}
}