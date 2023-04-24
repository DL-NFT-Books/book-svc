package middlewares

import (
	"net/http"

	"github.com/dl-nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CheckAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		address, err := helpers.ValidateJwt(r)
		if err != nil {
			ape.RenderErr(w, problems.Unauthorized())
			return
		}
		helpers.CtxUserAddress(address)
		next.ServeHTTP(w, r)
	})
}
