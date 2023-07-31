package comment

import (
	"context"
	"log"
)

type repository interface {
	writeCommentRepo
}

type writeCommentRepo interface {
	createComment(ctx context.Context, param Comment) (err error, comment Comment)
	getComment(ctx context.Context, commentId int) (err error, comment Comment, username string)
	getCommentList(ctx context.Context, photoId int) (err error, commentList []Comment)
	editComment(ctx context.Context, param Comment) (err error, comment Comment)
	removeComment(ctx context.Context, commentId int, userId int) (err error)
}

type commentService struct {
	repo repository
	r    commentRepo
}

func newCommentService(repo repository, r commentRepo) commentService {
	return commentService{
		repo: repo,
		r:    r,
	}
}

func (u commentService) createComment(ctx context.Context, req CommentReq) (err error, resp CommentResp) {
	var comment Comment

	err, comment = u.repo.createComment(ctx, req.CommentReqIntoComment())
	if err != nil {
		log.Println(err)
		return
	}

	if comment.CreatedAt == "" {
		return
	} else {
		resp = comment.CommentIntoCommentResp()
	}

	return err, resp
}

func (u commentService) getComment(ctx context.Context, req CommentReq) (err error, resp CommentResp) {
	err, comment, username := u.repo.getComment(ctx, req.CommentId)
	resp = comment.CommentIntoCommentResp()
	resp.Username = username

	return err, resp
}

func (u commentService) getCommentList(ctx context.Context, req CommentReq) (err error, respList []CommentResp) {
	err, commentList := u.repo.getCommentList(ctx, req.PhotoId)
	_ = commentList

	for _, comment := range commentList {
		respList = append(respList, comment.CommentIntoCommentResp())
	}

	return err, respList
}

func (u commentService) editComment(ctx context.Context, req CommentReq) (err error, resp CommentResp) {
	var comment Comment

	err, comment = u.repo.editComment(ctx, req.CommentReqIntoComment())
	if err != nil {
		log.Println(err)
		return
	}

	if comment.UpdatedAt == "" {
		return
	} else {
		resp = comment.CommentIntoCommentResp()
	}

	return err, resp
}

func (u commentService) removeComment(ctx context.Context, req CommentReq) (err error) {
	err = u.repo.removeComment(ctx, req.CommentId, req.UserId)

	return
}
