package api

import (
	"errors"
	"net/http"

	db "github.com/BinayRajbanshi/GoBasicBank/db/sqlc"
	"github.com/BinayRajbanshi/GoBasicBank/token"
	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBind(&req); err != nil { // meaning client has sent invalid data
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// .(*token.Payload) => this is called type assertion.  asserting that the value retrieved from the context is of type *token.Payload
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)

	if err != nil {
		// check what type of error I am getting from the database so that internal server error is not thrown randomly and appropriate status is thrown
		errCode := util.ErrorCode(err)
		if errCode == util.ForeignKeyViolation || errCode == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		util.ErrorCode(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id"  binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil { //binding failed means validation failed. so
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := server.store.GetAccount(ctx, req.ID)

	if err != nil {
		if errors.Is(err, util.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// account := db.Account{}

	// check with the help of middleware that the requested account belongs to this user
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account does not belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageNo  int32 `form:"page_no" binding:"required,min=1"`
	PerPage int32 `form:"per_page" binding:"required,min=1"`
}

func (server *Server) getAccounts(ctx *gin.Context) {
	var req listAccountRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	params := db.ListAccountsParams{
		Owner:  payload.Username,
		Limit:  req.PerPage,
		Offset: (req.PageNo - 1) * req.PerPage,
	}
	accounts, err := server.store.ListAccounts(ctx, params)

	if err != nil {
		if errors.Is(err, util.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
