package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/BinayRajbanshi/GoBasicBank/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// gin.HandlerFunc is just a function that takes in a *gin.Context and does something (like modify it, halt the request, or let it continue. In this case it sets the value of payload).
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	// returning gin.HandlerFunc as expected by authMiddleware
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		// case when authorization header is not provided
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// retrive the header and check if it has two parts (type token)
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// verify that the first part is of type bearer
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// extract the second part and verify the token
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// userinfo i.e. payload is sent through the context
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
