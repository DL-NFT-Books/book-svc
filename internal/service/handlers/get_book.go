package handlers

import (
	"net/http"

	"gitlab.com/tokend/nft-books/book-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetBookByIDRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	book, err := helpers.GetBookByID(r, req.ID)
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}
	if book == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	media, err := helpers.UnmarshalMedia(book.Banner, book.File)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	data, err := helpers.NewBook(book)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	media[0].Key = resources.NewKeyInt64(book.ID, resources.BANNER)
	media[1].Key = resources.NewKeyInt64(book.ID, resources.FILE)

	included := resources.Included{}
	included.Add(&media[0], &media[1])

	ape.Render(w, resources.BookResponse{
		Data:     data,
		Included: included,
	})
}
