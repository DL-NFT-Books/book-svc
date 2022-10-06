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

func CreateBook(w http.ResponseWriter, r *http.Request) {
	//TODO:check auth

	log := helpers.Log(r)

	req, err := requests.NewCreateBookRequest(r)
	if err != nil {
		log.WithError(err).Info("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	media := helpers.MarshalMedia(req.Attributes.Banner, req.Attributes.File)
	if media == nil {
		log.WithError(err).Info("failed to marshal media")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	book := data.Book{
		Title:       req.Attributes.Title,
		Description: req.Attributes.Description,
		Price:       req.Attributes.Price,
		Banner:      media[0],
		File:        media[1],
	}

	bookID, err := helpers.BooksQ(r).Insert(book)
	if err != nil {
		log.WithError(err).Info("failed to save book")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	req.Key = resources.NewKeyInt64(bookID, resources.BOOK)

	ape.Render(w, resources.BookResponse{
		Data: req,
	})
}

func newBook(book data.Book) (resources.Book, error) {
	banner, err := helpers.UnmarshalMedia(book.Banner)
	if err != nil {
		return resources.Book{}, err
	}

	file, err := helpers.UnmarshalMedia(book.File)
	if err != nil {
		return resources.Book{}, err
	}

	res := resources.Book{
		Key: resources.NewKeyInt64(book.ID, resources.BOOK),
		Attributes: resources.BookAttributes{
			Title:       book.Title,
			Description: book.Description,
			Price:       book.Price,
			Banner:      banner,
			File:        file,
		},
	}

	return res, nil
}
