package security

import (
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//extract token
		tokenString := request.Header.Get("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		//Validation and extract claims from token
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			http.Error(writer, "Unauthrithed", http.StatusUnauthorized)
			return
		}
		//inject claims into context request
		ctx := context.WithValue(request.Context(), "userClaims", claims)
		//toss to next handler
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
