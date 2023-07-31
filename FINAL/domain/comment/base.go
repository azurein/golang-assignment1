package comment

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.RouterGroup, db *sql.DB) {
	repo := newCommentRepo(db)
	svc := newCommentService(repo, repo)
	handler := newCommentHandler(svc)

	comment := router.Group("/comment")
	{
		comment.POST("", handler.createComment)
		comment.GET("", handler.getComment)
		comment.GET("list", handler.getCommentList)
		comment.PUT("", handler.editComment)
		comment.DELETE("", handler.removeComment)
	}
}
