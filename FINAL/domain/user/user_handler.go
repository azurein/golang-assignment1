package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	svc userService
}

func newUserHandler(svc userService) userHandler {
	return userHandler{
		svc: svc,
	}
}

func writeMessageError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"success": false,
		"error":   err,
	})
}

func (u userHandler) createUser(ctx *gin.Context) {
	var req = UserReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	err, resp := u.svc.createUser(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "user registration success",
	})
}

func (u userHandler) userLogin(ctx *gin.Context) {
	var req = UserReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	err, resp := u.svc.userLogin(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "user login success",
	})
}
