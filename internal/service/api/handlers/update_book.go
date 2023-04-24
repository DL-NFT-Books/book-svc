package handlers

import (
	"errors"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"

	"github.com/dl-nft-books/book-svc/internal/data"
	"github.com/dl-nft-books/book-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {

	logger := helpers.Log(r)
	networker := helpers.Networker(r)
	request, err := requests.NewUpdateBookRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	address := helpers.UserAddress(r)
	bookData, err := helpers.GetBookByID(r, requests.GetBookByIDRequest{
		ID: request.ID,
	})
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if bookData == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	book, err := helpers.NewBook(bookData)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	for _, net := range book.Attributes.Networks {
		network, err := networker.GetNetworkDetailedByChainID(net.Attributes.ChainId)
		if err != nil {
			logger.WithError(err).Error("default failed to check if network exists")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if network == nil {
			logger.WithError(err).Error("network does not exist")
			ape.RenderErr(w, problems.NotFound())
			return
		}
		isMarketplaceManager, err := helpers.CheckMarketplacePermission(*network, address)
		if err != nil {
			logger.WithError(err).Debug("failed to check is admin")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if !isMarketplaceManager {
			logger.WithFields(logan.F{"account": address}).Debug("you don't have permission to create book")
			ape.RenderErr(w, problems.Forbidden())
			return
		}
	}
	updateParams := data.BookUpdateParams{
		Banner:      nil,
		File:        nil,
		Description: nil,
	}

	// Collecting update params
	banner := request.Data.Attributes.Banner
	if banner != nil {
		if err = helpers.CheckBannerMimeType(banner.Attributes.MimeType, r); err != nil {
			helpers.Log(r).WithError(err).Error("failed to validate banner mime type")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if err = helpers.SetMediaLink(r, banner); err != nil {
			helpers.Log(r).WithError(err).Error("failed to set banner link")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		bannerMediaRaw := helpers.MarshalMedia(banner)
		updateParams.Banner = &bannerMediaRaw[0]
	}

	file := request.Data.Attributes.File
	if file != nil {
		if err = helpers.CheckFileMimeType(file.Attributes.MimeType, r); err != nil {
			helpers.Log(r).WithError(err).Error("failed to validate file mime type")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if err = helpers.SetMediaLink(r, file); err != nil {
			helpers.Log(r).WithError(err).Error("failed to set file link")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		fileMediaRaw := helpers.MarshalMedia(file)
		updateParams.File = &fileMediaRaw[0]
	}

	description := request.Data.Attributes.Description
	if description != nil {
		if len(*description) > requests.MaxDescriptionLength {
			err = errors.New(fmt.Sprintf("invalid description length (max len is %v)", requests.MaxDescriptionLength))
			helpers.Log(r).WithError(err).Error("failed to validate book description")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		updateParams.Description = description
	}

	// Updating collected params
	if err = helpers.DB(r).Books().Update(updateParams, request.ID); err != nil {
		helpers.Log(r).WithError(err).Error("failed to update book params")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, http.StatusNoContent)
}
