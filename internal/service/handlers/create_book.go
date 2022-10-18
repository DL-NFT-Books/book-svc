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
	req, err := requests.NewCreateBookRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = helpers.CheckMediaTypes(r, req.Banner.Attributes.MimeType, req.File.Attributes.MimeType)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	media := helpers.MarshalMedia(req.Banner, req.File)
	if media == nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	book := data.Book{
		Title:       req.Data.Attributes.Title,
		Description: req.Data.Attributes.Description,
		Price:       req.Data.Attributes.Price,
		Banner:      media[0],
		File:        media[1],
	}

	bookID, err := helpers.BooksQ(r).Insert(book)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	req.Data.Key = resources.NewKeyInt64(bookID, resources.BOOKS)
	req.Banner.Key = resources.NewKeyInt64(bookID, resources.BANNER)
	req.Data.Relationships.Banner.Data = &req.Banner.Key
	req.File.Key = resources.NewKeyInt64(bookID, resources.FILE)
	req.Data.Relationships.File.Data = &req.File.Key

	included := resources.Included{}
	included.Add(req.Banner, req.File)

	ape.Render(w, resources.BookResponse{
		Data:     req.Data,
		Included: included,
	})
}

func newBook(book *data.Book) (resources.Book, error) {
	if book == nil {
		return resources.Book{}, nil
	}

	bannerKey := resources.NewKeyInt64(book.ID, resources.BANNER)
	documentKey := resources.NewKeyInt64(book.ID, resources.FILE)

	res := resources.Book{
		Key: resources.NewKeyInt64(book.ID, resources.BOOKS),
		Attributes: resources.BookAttributes{
			Title:       book.Title,
			Description: book.Description,
			Price:       book.Price,
		},
		Relationships: resources.BookRelationships{
			Banner: resources.Relation{
				Data: &bannerKey,
			},
			File: resources.Relation{
				Data: &documentKey,
			},
		},
	}

	return res, nil
}
