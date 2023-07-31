package socialmedia

import (
	"finalproject/domain/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type socialmediaHandler struct {
	svc socialmediaService
}

func newSocialmediaHandler(svc socialmediaService) socialmediaHandler {
	return socialmediaHandler{
		svc: svc,
	}
}

func writeMessageError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"success": false,
		"error":   err,
	})
}

func (u socialmediaHandler) createSocialmedia(ctx *gin.Context) {
	var req = SocialmediaReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	accessToken := ctx.GetHeader("accessToken")
	err, userId := user.VerifyJWT(accessToken)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}
	req.UserId = int(userId)

	err, resp := u.svc.createSocialmedia(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "create socialmedia success",
	})
}

func (u socialmediaHandler) getSocialmedia(ctx *gin.Context) {
	var req = SocialmediaReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	err, resp := u.svc.getSocialmedia(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "get socialmedia success",
	})
}

func (u socialmediaHandler) getSocialmediaList(ctx *gin.Context) {
	err, resp := u.svc.getSocialmediaList(ctx.Request.Context())
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "get socialmedia list success",
	})
}

func (u socialmediaHandler) editSocialmedia(ctx *gin.Context) {
	var req = SocialmediaReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	accessToken := ctx.GetHeader("accessToken")
	err, userId := user.VerifyJWT(accessToken)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}
	req.UserId = int(userId)

	err, resp := u.svc.editSocialmedia(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "create socialmedia success",
	})
}

func (u socialmediaHandler) removeSocialmedia(ctx *gin.Context) {
	var req = SocialmediaReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	accessToken := ctx.GetHeader("accessToken")
	err, userId := user.VerifyJWT(accessToken)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}
	req.UserId = int(userId)

	err = u.svc.removeSocialmedia(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "remove socialmedia success",
	})
}
