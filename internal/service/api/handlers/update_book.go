package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
)

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewUpdateBookRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	book, err := helpers.GetBookByID(r, req.ID)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if book == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	banner := req.Data.Attributes.Banner
	if banner != nil {
		if err = helpers.CheckBannerMimeType(banner.Attributes.MimeType, r); err != nil {
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if err = helpers.SetMediaLink(r, banner); err != nil {
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		bannerMediaRaw := helpers.MarshalMedia(banner)

		if err = helpers.BooksQ(r).UpdateBanner(bannerMediaRaw[0], book.ID); err != nil {
			helpers.Log(r).WithError(err).Debug("failed to update book banner")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	file := req.Data.Attributes.File
	if file != nil {
		if err = helpers.CheckFileMimeType(file.Attributes.MimeType, r); err != nil {
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if err = helpers.SetMediaLink(r, file); err != nil {
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		fileMediaRaw := helpers.MarshalMedia(file)

		if err = helpers.BooksQ(r).UpdateFile(fileMediaRaw[0], book.ID); err != nil {
			helpers.Log(r).WithError(err).Debug("failed to update book file")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	title := req.Data.Attributes.Title
	if title != nil {
		if len(*title) > requests.MaxTitleLength {
			ape.RenderErr(w, problems.BadRequest(
				errors.New(
					fmt.Sprintf("invalid title length (max len is %v)", requests.MaxTitleLength)))...)
			return
		}

		if err = helpers.BooksQ(r).UpdateTitle(*title, book.ID); err != nil {
			helpers.Log(r).WithError(err).Debug("failed to update book title")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	description := req.Data.Attributes.Description
	if description != nil {
		if len(*description) > requests.MaxDescriptionLength {
			ape.RenderErr(w, problems.BadRequest(
				errors.New(
					fmt.Sprintf("invalid description length (max len is %v)", requests.MaxDescriptionLength)))...)
			return
		}

		if err = helpers.BooksQ(r).UpdateDescription(*description, book.ID); err != nil {
			helpers.Log(r).WithError(err).Debug("failed to update book description")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	ape.Render(w, http.StatusNoContent)
}
