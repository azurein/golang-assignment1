package user

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.RouterGroup, db *sql.DB) {
	repo := newUserRepo(db)
	svc := newUserService(repo, repo)
	handler := newUserHandler(svc)

	user := router.Group("/user")
	{
		user.POST("register", handler.createUser)
		user.POST("login", handler.userLogin)
	}
}
