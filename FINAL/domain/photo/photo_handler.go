package photo

import (
	"log"
	"net/http"

	"finalproject/domain/user"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	svc photoService
}

func newPhotoHandler(svc photoService) photoHandler {
	return photoHandler{
		svc: svc,
	}
}

func writeMessageError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"success": false,
		"error":   err,
	})
}

func (u photoHandler) createPhoto(ctx *gin.Context) {
	var req = PhotoReq{}
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

	err, resp := u.svc.createPhoto(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "create photo success",
	})
}

func (u photoHandler) getPhoto(ctx *gin.Context) {
	var req = PhotoReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	err, resp := u.svc.getPhoto(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "get photo success",
	})
}

func (u photoHandler) getPhotoList(ctx *gin.Context) {
	err, resp := u.svc.getPhotoList(ctx.Request.Context())
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "get photo list success",
	})
}

func (u photoHandler) editPhoto(ctx *gin.Context) {
	var req = PhotoReq{}

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

	err, resp := u.svc.editPhoto(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "create photo success",
	})
}

func (u photoHandler) removePhoto(ctx *gin.Context) {
	var req = PhotoReq{}

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

	err = u.svc.removePhoto(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "remove photo success",
	})
}
