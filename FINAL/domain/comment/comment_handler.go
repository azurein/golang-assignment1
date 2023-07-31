package comment

import (
	"finalproject/domain/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	svc commentService
}

func newCommentHandler(svc commentService) commentHandler {
	return commentHandler{
		svc: svc,
	}
}

func writeMessageError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"success": false,
		"error":   err,
	})
}

func (u commentHandler) createComment(ctx *gin.Context) {
	var req = CommentReq{}

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

	err, resp := u.svc.createComment(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "create comment success",
	})
}

func (u commentHandler) getComment(ctx *gin.Context) {
	var req = CommentReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	err, resp := u.svc.getComment(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "get comment success",
	})
}

func (u commentHandler) getCommentList(ctx *gin.Context) {
	var req = CommentReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	err, resp := u.svc.getCommentList(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "get comment list success",
	})
}

func (u commentHandler) editComment(ctx *gin.Context) {
	var req = CommentReq{}

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

	err, resp := u.svc.editComment(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "create comment success",
	})
}

func (u commentHandler) removeComment(ctx *gin.Context) {
	var req = CommentReq{}

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

	err = u.svc.removeComment(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	// TODO: refactor to fix err msg
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "remove comment success",
	})
}
