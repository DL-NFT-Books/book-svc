package handlers

import (
	"net/http"

	"gitlab.com/tokend/nft-books/book-svc/internal/data"
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

	books, err := helpers.BooksQ(r).Page(req.OffsetPageParams).Select()
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if books == nil {
		ape.Render(w, books)
		return
	}

	data, included, err := newBooksList(books)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.BookListResponse{
		Data:     data,
		Included: included,
	})
}

func newBooksList(books []data.Book) ([]resources.Book, resources.Included, error) {
	data := make([]resources.Book, len(books))
	included := resources.Included{}

	for i, book := range books {
		media, err := helpers.UnmarshalMedia(book.Banner, book.File)
		if err != nil {
			return nil, resources.Included{}, err
		}

		responseBook, err := newBook(&book)
		if err != nil {
			return nil, resources.Included{}, err
		}

		data[i] = responseBook

		media[0].Key = resources.NewKeyInt64(book.ID, resources.BANNER)
		media[1].Key = resources.NewKeyInt64(book.ID, resources.FILE)
		included.Add(&media[0], &media[1])
	}

	return data, included, nil
}
