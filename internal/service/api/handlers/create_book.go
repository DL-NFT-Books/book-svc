package handlers

import (
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
	"time"

	"github.com/dl-nft-books/book-svc/internal/data"
	"github.com/dl-nft-books/book-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/book-svc/internal/service/api/requests"
	"github.com/dl-nft-books/book-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	networker := helpers.Networker(r)

	request, err := requests.NewCreateBookRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	address := helpers.UserAddress(r)
	for _, net := range request.Data.Attributes.Networks {
		network, err := networker.GetNetworkDetailedByChainID(net.Attributes.ChainId)
		if err != nil {
			logger.WithError(err).Error("failed to check if network exists")
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
	// Validating info
	banner := request.Data.Attributes.Banner
	file := request.Data.Attributes.File

	if err = helpers.CheckMediaTypes(r, banner.Attributes.MimeType, file.Attributes.MimeType); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Setting banner link
	if err = helpers.SetMediaLink(r, &banner); err != nil {
		logger.WithError(err).Error("failed to set banner link")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Setting file link
	if err = helpers.SetMediaLink(r, &file); err != nil {
		logger.WithError(err).Error("failed to set file link")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	media := helpers.MarshalMedia(&banner, &file)
	if media == nil {
		logger.Error("failed to marshal media")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if err != nil {
		logger.WithError(err).Error("failed to get last token id")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// Saving book to the database
	book := data.Book{
		Description: request.Data.Attributes.Description,
		CreatedAt:   time.Now(),
		Banner:      media[0],
		File:        media[1],
	}
	db := helpers.DB(r)
	var bookId int64

	if err = db.Transaction(func() error {
		// Inserting book
		bookId, err = db.Books().Insert(book)
		if err != nil {
			return errors.Wrap(err, "failed to save book")
		}

		// Inserting book networks

		var bookNetwork []data.BookNetwork
		for _, network := range request.Data.Attributes.Networks {
			bookNetwork = append(bookNetwork, data.BookNetwork{
				BookId:          bookId,
				ContractAddress: network.Attributes.ContractAddress,
				ChainId:         network.Attributes.ChainId,
			})
		}
		err = db.Books().InsertNetwork(bookNetwork...)

		return nil
	},
	); err != nil {
		logger.WithError(err).Error("failed to execute insertion tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, resources.KeyResponse{
		Data: resources.NewKeyInt64(bookId, resources.BOOKS),
	})
}
