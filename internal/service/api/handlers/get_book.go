package handlers

import (
	"net/http"

	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	req, err := requests.NewGetBookByIDRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	book, err := helpers.GetBookByID(r, req.ID)
	if err != nil {
		logger.WithError(err).Error("failed to get book by id")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if book == nil {
		logger.Error("book not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	data, err := helpers.NewBook(book)
	if err != nil {
		logger.WithError(err).Error("failed to form up book response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.BookResponse{
		Data: *data,
	})
}
