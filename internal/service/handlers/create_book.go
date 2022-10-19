package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"
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
		Title:           req.Data.Attributes.Title,
		Description:     req.Data.Attributes.Description,
		Price:           req.Data.Attributes.Price,
		ContractAddress: req.Data.Attributes.ContractAddress,
		ContractName:    req.Data.Attributes.ContractName,
		ContractVersion: req.Data.Attributes.ContractVersion,
		Banner:          media[0],
		File:            media[1],
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
