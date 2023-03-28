package handlers

import (
	"github.com/dl-nft-books/book-svc/solidity/generated/rolemanager"
	"github.com/ethereum/go-ethereum/common"
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
	if _, err := networker.GetNetworkDetailedByChainID(request.Data.Attributes.ChainId); err != nil {
		logger.WithError(err).Error("default failed to check if network exists")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	network, err := networker.GetNetworkDetailedByChainID(request.Data.Attributes.ChainId)

	address := r.Context().Value("address").(string)

	roleManager, err := rolemanager.NewRolemanager(common.HexToAddress(network.FactoryAddress), network.RpcUrl)
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
		logger.Debug("not admin's address")
		ape.RenderErr(w, problems.Forbidden())
		return
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
		Description:     request.Data.Attributes.Description,
		CreatedAt:       time.Now(),
		ContractAddress: "mocked",
		Banner:          media[0],
		File:            media[1],
		TokenId:         tokenContractId,
		DeployStatus:    resources.DeployPending,
		ChainId:         request.Data.Attributes.ChainId,
	}

	db := helpers.DB(r)
	var bookId int64

	if err = db.Transaction(func() error {
		// Inserting book
		bookId, err = db.Books().Insert(book)
		if err != nil {
			return errors.Wrap(err, "failed to save book")
		}

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
