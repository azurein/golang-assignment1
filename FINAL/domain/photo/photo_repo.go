package photo

import (
	"context"
	"database/sql"
	"log"
)

type photoRepo struct {
	db *sql.DB
}

func newPhotoRepo(db *sql.DB) photoRepo {
	return photoRepo{
		db: db,
	}
}

func (u photoRepo) createPhoto(ctx context.Context, photo Photo) (err error, _ Photo) {
	// TODO: add check user exists
	query := `
		INSERT INTO "photo" (user_id, title, caption, photo_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, now(), now())
		RETURNING photo_id, created_at, updated_at
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, photo
	}
	defer stmt.Close()

	var newPhotoId int
	var newPhotoCreatedAt, newPhotoUpdatedAt string
	err = stmt.QueryRow(photo.UserId, photo.Title, photo.Caption, photo.PhotoUrl).Scan(
		&newPhotoId, &newPhotoCreatedAt, &newPhotoUpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return err, photo
	}
	photo.PhotoId = newPhotoId
	photo.CreatedAt = newPhotoCreatedAt
	photo.UpdatedAt = newPhotoUpdatedAt

	return err, photo
}

func (u photoRepo) getPhoto(ctx context.Context, photoId int) (err error, photo Photo, username string) {
	query := `
		SELECT p.photo_id, p.user_id, u.username, p.title, p.caption, p.photo_url, p.created_at, p.updated_at
		FROM "photo" as p
		JOIN "user" as u
		ON p.user_id = u.user_id 
		WHERE p.deleted = false
		AND u.deleted = false
		AND p.photo_id = $1
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, photo, username
	}
	defer stmt.Close()

	row, err := stmt.QueryContext(ctx, photoId)
	if err != nil {
		log.Println(err)
		return err, photo, username
	}
	if row.Next() {
		err := row.Scan(
			&photo.PhotoId,
			&photo.UserId,
			&username,
			&photo.Title,
			&photo.Caption,
			&photo.PhotoUrl,
			&photo.CreatedAt,
			&photo.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return err, photo, username
		}
	}

	return err, photo, username
}

func (u photoRepo) getPhotoList(ctx context.Context) (err error, photoList []Photo) {
	// TODO: select username
	query := `
		SELECT p.photo_id, p.user_id, p.title, p.caption, p.photo_url, p.created_at, p.updated_at
		FROM "photo" as p
		JOIN "user" as u
		ON p.user_id = u.user_id 
		WHERE p.deleted = false
		AND u.deleted = false
		ORDER BY p.updated_at DESC
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, photoList
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Println(err)
		return err, photoList
	}
	for rows.Next() {
		var photo Photo
		err := rows.Scan(
			&photo.PhotoId,
			&photo.UserId,
			&photo.Title,
			&photo.Caption,
			&photo.PhotoUrl,
			&photo.CreatedAt,
			&photo.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return err, photoList
		}
		photoList = append(photoList, photo)
	}

	return err, photoList
}

func (u photoRepo) editPhoto(ctx context.Context, photo Photo) (err error, _ Photo) {
	query := `
		UPDATE "photo" SET 
		title = $1, 
		caption = $2, 
		photo_url = $3, 
		updated_at = now()
		WHERE photo_id = $4
		AND user_id = $5
		RETURNING created_at, updated_at
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, photo
	}
	defer stmt.Close()

	var newPhotoCreatedAt, newPhotoUpdatedAt string
	err = stmt.QueryRow(photo.Title, photo.Caption, photo.PhotoUrl, photo.PhotoId, photo.UserId).Scan(
		&newPhotoCreatedAt, &newPhotoUpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return err, photo
	}
	photo.CreatedAt = newPhotoCreatedAt
	photo.UpdatedAt = newPhotoUpdatedAt

	return err, photo
}

func (u photoRepo) removePhoto(ctx context.Context, photoId int, userId int) (err error) {
	query := `
		UPDATE "photo" SET updated_at = now(), deleted = true
		WHERE photo_id = $1 AND user_id = $2
	`
	stmt, err := u.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return
	}

	_, err = stmt.ExecContext(ctx, photoId, userId)
	if err != nil {
		log.Println(err)
		return err
	}

	return
}
