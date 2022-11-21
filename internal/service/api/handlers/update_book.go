package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
)

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

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

	updateParams := data.BookUpdateParams{}

	// collecting update params
	banner := req.Data.Attributes.Banner
	if banner != nil {
		if err = helpers.CheckBannerMimeType(banner.Attributes.MimeType, r); err != nil {
			logger.WithError(err).Error("failed to validate banner mime type")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if err = helpers.SetMediaLink(r, banner); err != nil {
			logger.WithError(err).Error("failed to set banner link")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		bannerMediaRaw := helpers.MarshalMedia(banner)
		updateParams.Banner = &bannerMediaRaw[0]
	}

	file := req.Data.Attributes.File
	if file != nil {
		if err = helpers.CheckFileMimeType(file.Attributes.MimeType, r); err != nil {
			logger.WithError(err).Error("failed to validate file mime type")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if err = helpers.SetMediaLink(r, file); err != nil {
			logger.WithError(err).Error("failed to set file link")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		fileMediaRaw := helpers.MarshalMedia(file)
		updateParams.File = &fileMediaRaw[0]
	}

	title := req.Data.Attributes.Title
	if title != nil {
		if len(*title) > requests.MaxTitleLength {
			err = errors.New(fmt.Sprintf("invalid title length (max len is %v)", requests.MaxTitleLength))
			logger.WithError(err).Error("failed to validate book title")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		updateParams.Title = title
	}

	description := req.Data.Attributes.Description
	if description != nil {
		if len(*description) > requests.MaxDescriptionLength {
			err = errors.New(fmt.Sprintf("invalid description length (max len is %v)", requests.MaxDescriptionLength))
			logger.WithError(err).Error("failed to validate book description")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		updateParams.Description = description
	}

	// updating collected params
	if err = helpers.DB(r).Books().Update(updateParams, req.ID); err != nil {
		logger.WithError(err).Error("failed to update book params")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, http.StatusNoContent)
}
