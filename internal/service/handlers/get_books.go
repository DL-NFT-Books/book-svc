package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetBooksRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	books, err := helpers.BooksQ(r).Page(req.OffsetPageParams).Select()
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if books == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	response, err := newBooksList(books)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.BookListResponse{
		Data: response,
	})
}

func newBooksList(books []data.Book) ([]resources.Book, error) {
	res := make([]resources.Book, len(books))
	for i, book := range books {
		responseBook, err := newBook(book)
		if err != nil {
			return nil, err
		}
		res[i] = responseBook
	}

	return res, nil
}
