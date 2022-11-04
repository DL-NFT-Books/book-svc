package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/requests"
)

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewUpdateBookRequest(r)
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

	err = helpers.CheckMediaTypes(r, req.Data.Attributes.Banner.Attributes.MimeType, req.Data.Attributes.File.Attributes.MimeType)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	media := helpers.MarshalMedia(&req.Data.Attributes.Banner, &req.Data.Attributes.File)
	if media == nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	bookToUpdate := data.Book{
		ID:          req.ID,
		Title:       req.Data.Attributes.Title,
		Description: req.Data.Attributes.Description,
		Banner:      media[0],
		File:        media[1],
	}

	err = helpers.BooksQ(r).Update(bookToUpdate)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, http.StatusNoContent)
}
