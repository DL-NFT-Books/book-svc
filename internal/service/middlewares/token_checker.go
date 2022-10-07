package middlewares

import (
	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/helpers"
	"net/http"
	"strings"
	"time"
)

const (
	TokenHeader  = "Authorization"
	RequiredRole = "admin"
)

type AccessTokenClaims struct {
	Role  string `json:"role"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CheckAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tokenParam := request.Header.Get(TokenHeader)
		if tokenParam == "" {
			ape.RenderErr(writer, problems.Unauthorized())
			return
		}

		values := strings.Split(tokenParam, "Bearer ")
		if len(values) != 2 {
			ape.RenderErr(writer, problems.Unauthorized())
			return
		}
		tokenString := values[1]

		err := ValidateJWT(tokenString, request)
		if err != nil {
			ape.RenderErr(writer, problems.Unauthorized())
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func ValidateJWT(tokenString string, r *http.Request) error {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(helpers.JWT(r).SignatureKey), nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}

	return checkClaims(claims)
}

func checkClaims(claims *AccessTokenClaims) error {
	if claims.ExpiresAt.Before(time.Now()) {
		return errors.New("expired access token")
	}

	if claims.Role != RequiredRole {
		return errors.New("invalid user role")
	}

	return nil
}
