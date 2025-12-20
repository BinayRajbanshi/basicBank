package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/BinayRajbanshi/GoBasicBank/token"
	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/gin-gonic/gin"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	// accpet and bind refresh token
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// verify token's validity
	payload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// check correct token type
	if payload.TokenType != token.TokenTypeRefreshToken {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Invalid Refresh token")))
		return
	}

	// obtain token from session db
	session, err := server.store.GetSession(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, util.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New("No previous session created so access token generation unavailable. ")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// check if blocked or not
	if session.IsBlocked {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Session Blocked")))
		return
	}

	// check if sent by correct user or not
	if session.Username != payload.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Invalid Username"))) // what if the user changes the username ?
		return
	}
	// compare refresh token from session db and decoded payload
	if session.RefreshToken != req.RefreshToken {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Session Mismatched")))
		return
	}

	// finally check token's expiration
	if time.Now().After(session.ExpiresAt) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Session Expired. Please Relogin")))
		return
	}

	// if all of the conditions are fulfilled send back newly generated access token with short life time
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(payload.Username, server.config.AccessTokenDuration, token.TokenTypeAccessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res := &renewAccessTokenResponse{
		AccessToken: accessToken,
		ExpiresAt:   accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, res)
}
