package handlers

import (
	"math/big"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/responses"
	"gitlab.com/tokend/nft-books/book-svc/internal/signature"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

const tokenIdIncrementKey = "token_id_increment"

func CreateBook(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	req, err := requests.NewCreateBookRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	banner := req.Data.Attributes.Banner
	file := req.Data.Attributes.File

	if err = helpers.CheckMediaTypes(r, banner.Attributes.MimeType, file.Attributes.MimeType); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if err = helpers.SetMediaLinks(r, &banner, &file); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	media := helpers.MarshalMedia(&banner, &file)
	if media == nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tokenPrice, ok := big.NewInt(0).SetString(req.Data.Attributes.Price, 10)
	if !ok {
		logger.Error("failed to cast price to big.Int")
		ape.RenderErr(w, problems.InternalError())
	}

	lastTokenContractID, err := helpers.GenerateTokenID(r)
	if err != nil {
		logger.WithError(err).Debug("failed to generate token id")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// forming signature createInfo
	config := helpers.DeploySignatureConfig(r)

	domainData := signature.EIP712DomainData{
		VerifyingAddress: config.TokenFactoryAddress,
		ContractName:     config.TokenFactoryName,
		ContractVersion:  config.TokenFactoryVersion,
		ChainID:          config.ChainId,
	}

	createInfo := signature.CreateInfo{
		TokenContractId:  lastTokenContractID + 1,
		TokenName:        req.Data.Attributes.TokenName,
		TokenSymbol:      req.Data.Attributes.TokenSymbol,
		PricePerOneToken: tokenPrice,
	}

	// signing
	signature, err := signature.SignCreateInfo(&createInfo, &domainData, config)
	if err != nil {
		logger.WithError(err).Debug("failed to generate eip712 create signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// saving book
	book := data.Book{
		Title:           req.Data.Attributes.Title,
		Description:     req.Data.Attributes.Description,
		CreatedAt:       time.Now(),
		Price:           req.Data.Attributes.Price,
		ContractAddress: "mocked",
		ContractName:    "mocked",
		ContractSymbol:  "mocked",
		ContractVersion: config.TokenFactoryVersion,
		Banner:          media[0],
		File:            media[1],
		Deleted:         false,
		TokenId:         createInfo.TokenContractId,
		DeployStatus:    resources.DeployPending,
		LastBlock:       0,
	}

	bookId, err := helpers.BooksQ(r).Insert(book)
	if err != nil {
		logger.WithError(err).Debug("failed to save book")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// update kv if operation was ended successfully
	if err = helpers.KeyValueQ(r).Upsert(data.KeyValue{
		Key:   tokenIdIncrementKey,
		Value: strconv.FormatInt(createInfo.TokenContractId, 10),
	}); err != nil {
		logger.WithError(err).Debug("failed to update last created token id")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewSignCreateResponse(bookId, createInfo.TokenContractId, *signature))
}
