package handlers

import (
	"net/http"

	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	req, err := requests.NewGetBooksRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	books, err := helpers.GetBookListByRequest(r, &req)
	if err != nil {
		logger.WithError(err).Error("failed to get books")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if books == nil {
		ape.Render(w, resources.BookListResponse{
			Data: make([]resources.Book, 0),
		})
		return
	}

	data, err := helpers.NewBooksList(books)
	if err != nil {
		logger.WithError(err).Error("failed to form up book list response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.BookListResponse{
		Data: data,
	})
}
