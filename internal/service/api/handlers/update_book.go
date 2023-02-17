package handlers

import (
	"errors"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
)

var invalidContractNameErr = errors.New("invalid contract name length")

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateBookRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	book, err := helpers.GetBookByID(r, request.ID)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if book == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	updateParams := data.BookUpdateParams{
		Contract:           request.Data.Attributes.ContractAddress,
		DeployStatus:       request.Data.Attributes.DeployStatus,
		Price:              request.Data.Attributes.Price,
		Symbol:             request.Data.Attributes.TokenSymbol,
		VoucherToken:       request.Data.Attributes.VoucherToken,
		VoucherTokenAmount: request.Data.Attributes.VoucherTokenAmount,
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

	title := request.Data.Attributes.Title
	if title != nil {
		if len(*title) > requests.MaxTitleLength {
			err = errors.New(fmt.Sprintf("invalid title length (max len is %v)", requests.MaxTitleLength))
			helpers.Log(r).WithError(err).Error("failed to validate book's title")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		updateParams.Title = title
	}

	contractName := request.Data.Attributes.ContractName
	if contractName != nil {
		if len(*contractName) > requests.MaxTitleLength {
			helpers.Log(r).WithFields(logan.F{"max_title_len": requests.MaxTitleLength}).Error(invalidContractNameErr)
			ape.RenderErr(w, problems.BadRequest(invalidContractNameErr)...)
			return
		}

		updateParams.ContractName = contractName
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
