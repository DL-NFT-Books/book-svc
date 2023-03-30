package handlers

import (
	"github.com/dl-nft-books/book-svc/solidity/generated/contractsregistry"
	"github.com/dl-nft-books/book-svc/solidity/generated/rolemanager"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
	"strconv"
	"time"

	"github.com/dl-nft-books/book-svc/internal/data"
	"github.com/dl-nft-books/book-svc/internal/data/postgres"
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

	address := r.Context().Value("address").(string)
	for _, chainID := range request.Data.Attributes.ChainIds {
		network, err := networker.GetNetworkDetailedByChainID(chainID)
		if err != nil {
			logger.WithError(err).Error("default failed to check if network exists")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		contractRegistry, err := contractsregistry.NewContractsregistry(common.HexToAddress(network.FactoryAddress), network.RpcUrl)
		if err != nil {
			logger.WithError(err).Debug("failed to create contract registry")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		roleManagerContract, err := contractRegistry.GetRoleManagerContract(nil)
		if err != nil {
			logger.WithError(err).Debug("failed to get role manager contract")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		roleManager, err := rolemanager.NewRolemanager(roleManagerContract, network.RpcUrl)
		if err != nil {
			logger.WithError(err).Debug("failed to create role manager")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		isAdmin, err := roleManager.RolemanagerCaller.IsAdmin(nil, common.HexToAddress(address))
		if err != nil {
			logger.WithError(err).Debug("failed to check is admin")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if !isAdmin {
			isManager, err := roleManager.RolemanagerCaller.IsAdmin(nil, common.HexToAddress(address))
			if err != nil {
				logger.WithError(err).Debug("failed to check is admin")
				ape.RenderErr(w, problems.InternalError())
				return
			}
			if !isManager {
				logger.WithFields(logan.F{"account": address}).Debug("you don't have permission to create book")
				ape.RenderErr(w, problems.Forbidden())
				return
			}
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

	lastTokenContractID, err := helpers.GetLastTokenID(r)
	if err != nil {
		logger.WithError(err).Error("failed to get last token id")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tokenContractId := lastTokenContractID + 1
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
		for _, chainID := range request.Data.Attributes.ChainIds {
			bookNetwork = append(bookNetwork, data.BookNetwork{
				BookId:  bookId,
				TokenId: tokenContractId,
				ChainId: chainID,
			})
		}
		err = db.Books().InsertNetwork(bookNetwork...)
		// Updating last token id
		if err = db.KeyValue().Upsert(data.KeyValue{
			Key:   postgres.TokenIdIncrementKey,
			Value: strconv.FormatInt(tokenContractId, 10),
		}); err != nil {
			return errors.Wrap(err, "failed to update last created token id")
		}

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
