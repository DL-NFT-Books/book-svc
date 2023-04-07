package handlers

import (
	"github.com/dl-nft-books/book-svc/internal/data"
	"github.com/dl-nft-books/book-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
)

func AddBookNetwork(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	networker := helpers.Networker(r)

	request, err := requests.NewAddBookNetworkRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	address := r.Context().Value("address").(string)
	for _, net := range request.Data {
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
		isMarketplaceManager, err := helpers.CheckMarketplacePerrmision(*network, address)
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
	bookData, err := helpers.DB(r).Books().FilterByID(request.BookId).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get book")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if bookData == nil {
		logger.WithError(err).Error("book with such id not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	book, err := helpers.NewBook(bookData)
	for _, network := range book.Attributes.Networks {
		for _, reqNetworks := range request.Data {
			if network.Attributes.ChainId == reqNetworks.Attributes.ChainId {
				logger.WithError(err).Error("book network is already exist")
				ape.RenderErr(w, problems.Conflict())
				return
			}
		}
	}
	var bookNetwork []data.BookNetwork
	for _, network := range request.Data {
		bookNetwork = append(bookNetwork, data.BookNetwork{
			BookId:          request.BookId,
			ContractAddress: network.Attributes.ContractAddress,
			ChainId:         network.Attributes.ChainId,
		})
	}
	if err = helpers.DB(r).Books().InsertNetwork(bookNetwork...); err != nil {
		logger.WithError(err).Error("failed to execute insertion tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, http.StatusNoContent)
}
