package comment

import (
	"context"
	"database/sql"
	"log"
)

type commentRepo struct {
	db *sql.DB
}

func newCommentRepo(db *sql.DB) commentRepo {
	return commentRepo{
		db: db,
	}
}

func (u commentRepo) createComment(ctx context.Context, comment Comment) (err error, _ Comment) {
	// TODO: add check user
	query := `SELECT photo_id FROM "photo" WHERE deleted = false AND photo_id = $1`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, comment
	}
	defer stmt.Close()

	row, err := stmt.QueryContext(ctx, comment.PhotoId)
	if err != nil {
		log.Println(err)
		return err, comment
	}
	if row.Next() {
		query := `
			INSERT INTO "comment" (photo_id, user_id, message, created_at, updated_at)
			VALUES ($1, $2, $3, now(), now())
			RETURNING comment_id, created_at, updated_at
		`
		stmt, err := u.db.Prepare(query)
		if err != nil {
			log.Println(err)
			return err, comment
		}

		var newCommentId int
		var newCommentCreatedAt, newCommentUpdatedAt string
		err = stmt.QueryRow(comment.PhotoId, comment.UserId, comment.Message).Scan(
			&newCommentId, &newCommentCreatedAt, &newCommentUpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return err, comment
		}
		comment.CommentId = newCommentId
		comment.CreatedAt = newCommentCreatedAt
		comment.UpdatedAt = newCommentUpdatedAt
	}

	return err, comment
}

func (u commentRepo) getComment(ctx context.Context, commentId int) (err error, comment Comment, username string) {
	query := `
		SELECT c.comment_id, c.photo_id, c.user_id, u.username, c.message, c.created_at, c.updated_at
		FROM "comment" as c
		JOIN "user" as u
		ON c.user_id = u.user_id 
		WHERE c.deleted = false
		AND u.deleted = false
		AND c.comment_id = $1
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, comment, username
	}
	defer stmt.Close()

	row, err := stmt.QueryContext(ctx, commentId)
	if err != nil {
		log.Println(err)
		return err, comment, username
	}
	if row.Next() {
		err := row.Scan(
			&comment.CommentId,
			&comment.PhotoId,
			&comment.UserId,
			&username,
			&comment.Message,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return err, comment, username
		}
	}

	return err, comment, username
}

func (u commentRepo) getCommentList(ctx context.Context, photoId int) (err error, commentList []Comment) {
	// TODO: select username
	query := `
		SELECT c.comment_id, c.photo_id, c.user_id, c.message, c.created_at, c.updated_at
		FROM "comment" as c
		JOIN "user" as u
		ON c.user_id = u.user_id 
		WHERE c.deleted = false
		AND u.deleted = false
		AND c.photo_id = $1
		ORDER BY c.updated_at DESC
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, commentList
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, photoId)
	if err != nil {
		log.Println(err)
		return err, commentList
	}
	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.CommentId,
			&comment.PhotoId,
			&comment.UserId,
			&comment.Message,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return err, commentList
		}
		commentList = append(commentList, comment)
	}

	return err, commentList
}

func (u commentRepo) editComment(ctx context.Context, comment Comment) (err error, _ Comment) {
	query := `SELECT photo_id FROM "photo" WHERE deleted = false AND photo_id = $1`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, comment
	}
	defer stmt.Close()

	row, err := stmt.QueryContext(ctx, comment.PhotoId)
	if err != nil {
		log.Println(err)
		return err, comment
	}
	if row.Next() {
		query := `
			UPDATE "comment" SET 
			message = $1, 
			updated_at = now()
			WHERE comment_id = $2
			AND user_id = $3
			RETURNING created_at, updated_at
		`
		stmt, err := u.db.Prepare(query)
		if err != nil {
			log.Println(err)
			return err, comment
		}
		defer stmt.Close()

		var newCommentCreatedAt, newCommentUpdatedAt string
		err = stmt.QueryRow(comment.Message, comment.CommentId, comment.UserId).Scan(
			&newCommentCreatedAt, &newCommentUpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return err, comment
		}
		comment.CreatedAt = newCommentCreatedAt
		comment.UpdatedAt = newCommentUpdatedAt
	}

	return err, comment
}

func (u commentRepo) removeComment(ctx context.Context, commentId int, userId int) (err error) {
	query := `
		UPDATE "comment" SET updated_at = now(), deleted = true
		WHERE comment_id = $1 AND user_id = $2
	`
	stmt, err := u.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return
	}

	_, err = stmt.ExecContext(ctx, commentId, userId)
	if err != nil {
		log.Println(err)
		return err
	}

	return
}
