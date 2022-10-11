package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/handlers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/middlewares"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxJWT(s.jwt),
			helpers.CtxBooksQ(postgres.NewBooksQ(s.db)),
			helpers.CtxMimeTypes(s.mimeTypes),
		),
		middlewares.CheckAccessToken,
	)
	r.Route("/integrations/books", func(r chi.Router) {
		r.Post("/", handlers.CreateBook)
		r.Get("/", handlers.GetBooks)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetBookByID)
		})
	})

	return r
}
