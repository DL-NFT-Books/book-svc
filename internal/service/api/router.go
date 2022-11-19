package api

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	handlers2 "gitlab.com/tokend/nft-books/book-svc/internal/service/api/handlers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/middlewares"

	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/postgres"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxBooksQ(postgres.NewBooksQ(s.db)),
			helpers.CtxKeyValueQ(postgres.NewKeyValueQ(s.db)),
			helpers.CtxMimeTypes(s.mimeTypes),
			helpers.CtxDeploySignature(s.deploySignatureConf),

			helpers.CtxDoormanConnector(cfg.DoormanConnector()),
			helpers.CtxDocumenterConnector(*cfg.DocumenterConnector()),
			helpers.CtxNetworkerConnector(*cfg.NetworkConnector()),
		),
	)

	r.Route("/integrations/books", func(r chi.Router) {
		r.With(middlewares.CheckAccessToken).
			Post("/", handlers2.CreateBook)

		r.Get("/", handlers2.GetBooks)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers2.GetBookByID)

			r.With(middlewares.CheckAccessToken).
				Patch("/", handlers2.UpdateBookByID)
			// TODO investigate
			//r.With(middlewares.CheckAccessToken).
			//	Delete("/", handlers.DeleteBookByID)
		})
	})

	return r
}
