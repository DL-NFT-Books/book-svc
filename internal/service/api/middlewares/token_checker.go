package middlewares

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
)

func CheckAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := helpers.ValidateJwt(r); err != nil {
			ape.RenderErr(w, problems.Unauthorized())
			return
		}

		next.ServeHTTP(w, r)
	})
}
