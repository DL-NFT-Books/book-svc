package handlers

import (
	"net/http"

	"gitlab.com/tokend/nft-books/book-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetBooksRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	books, err := helpers.GetBooksByPage(r, req.OffsetPageParams)
	if err != nil {
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
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.BookListResponse{
		Data: data,
	})
}
