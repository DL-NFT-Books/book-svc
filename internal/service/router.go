package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/handlers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/middlewares"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxBooksQ(postgres.NewBooksQ(s.db)),
			helpers.CtxMimeTypes(s.mimeTypes),
			helpers.CtxDoormanConnector(cfg.DoormanConnector()),
			helpers.CtxDocumenterConnector(*cfg.DocumenterConnector()),
		),
	)

	r.Route("/integrations/books", func(r chi.Router) {
		r.With(middlewares.CheckAccessToken).
			Post("/", handlers.CreateBook)

		r.Get("/", handlers.GetBooks)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetBookByID)

			r.With(middlewares.CheckAccessToken).
				Patch("/", handlers.UpdateBookByID)
			// TODO investigate
			//r.With(middlewares.CheckAccessToken).
			//	Delete("/", handlers.DeleteBookByID)
		})
	})

	return r
}
