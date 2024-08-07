package middlewares

import (
	"fmt"
	"github.com/Khvan-Group/common-library/utils"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"
	"net/http"
	"strconv"
	"strings"
)

func AuthMiddleware(next http.Handler, roles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isInternalService, err := strconv.ParseBool(r.Header.Get("X-Is-Internal-Service"))
		if err != nil {
			isInternalService = false
		}

		if isInternalService {
			next.ServeHTTP(w, r)
			return
		}

		token, err := verifyToken(r)
		if err != nil || !token.Valid {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		currentUserRole := utils.ToString(claims["role"])
		if len(roles) != 0 && roles[0] != "" && !utils.ContainsString(roles, currentUserRole) {
			http.Error(w, `{"error": "Доступ запрещен."}`, http.StatusForbidden)
			return
		}

		context.Set(r, "login", claims["iss"])
		context.Set(r, "role", claims["role"])
		next.ServeHTTP(w, r)
	})
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	bearerToken := r.Header.Get("Authorization")
	tokenString := ""
	if strings.HasPrefix(bearerToken, "Bearer ") {
		tokenString = strings.TrimPrefix(bearerToken, "Bearer ")
	}
	secretKey := []byte(utils.GetEnv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
