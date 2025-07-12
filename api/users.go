package api

import (
	"net/http"
	"time"

	db "github.com/BinayRajbanshi/GoBasicBank/db/sqlc"
	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Fullname string `json:"full_name" binding:"required"`
}

type UserPublicResponse struct {
	Username          string    `json:"username" binding:"required,alphanum"`
	Email             string    `json:"email" binding:"required,email"`
	Fullname          string    `json:"full_name" binding:"required"`
	CreatedAt         time.Time `json:"created_at" binding:"required"`
	PasswordChangedAt time.Time `json:"password_changed_at" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	args := db.CreateUserParams{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		FullName:       req.Fullname,
	}

	user, err := server.store.CreateUser(ctx, args)
	if err != nil {
		errCode := util.ErrorCode(err)
		if errCode == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := UserPublicResponse{
		Username:          user.Username,
		Email:             user.Email,
		Fullname:          user.FullName,
		CreatedAt:         user.CreatedAt,
		PasswordChangedAt: user.PasswordChangedAt,
	}
	ctx.JSON(http.StatusOK, response)

}
