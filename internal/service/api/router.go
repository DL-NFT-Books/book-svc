package api

import (
	"github.com/dl-nft-books/book-svc/internal/service/api/handlers"
	"github.com/dl-nft-books/book-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/book-svc/internal/service/api/middlewares"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"github.com/dl-nft-books/book-svc/internal/data/postgres"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			// Base configuration
			helpers.CtxLog(s.log),
			helpers.CtxDB(postgres.NewDB(s.db)),

			// Service configs
			helpers.CtxMimeTypes(s.mimeTypes),

			// Connectors
			helpers.CtxDoormanConnector(s.doorman),
			helpers.CtxDocumenterConnector(*s.documenter),
			helpers.CtxNetworkerConnector(*s.cfg.NetworkConnector()),
		),
	)

	r.Route("/integrations/books", func(r chi.Router) {
		r.With(middlewares.CheckAccessToken).Post("/", handlers.CreateBook)
		r.Get("/", handlers.ListBooks)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetBookByID)
			r.With(middlewares.CheckAccessToken).Patch("/", handlers.UpdateBookByID)
			r.With(middlewares.CheckAccessToken).Post("/network", handlers.AddBookNetwork)
		})
	})

	return r
}
