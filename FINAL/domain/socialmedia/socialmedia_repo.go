package socialmedia

import (
	"context"
	"database/sql"
	"log"
)

type socialmediaRepo struct {
	db *sql.DB
}

func newSocialmediaRepo(db *sql.DB) socialmediaRepo {
	return socialmediaRepo{
		db: db,
	}
}

func (u socialmediaRepo) createSocialmedia(ctx context.Context, socialmedia Socialmedia) (err error, _ Socialmedia) {
	// TODO: add check user
	query := `
		INSERT INTO "socialmedia" (user_id, name, socialmedia_url, created_at, updated_at)
		VALUES ($1, $2, $3, now(), now())
		RETURNING socialmedia_id, created_at, updated_at
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, socialmedia
	}
	defer stmt.Close()

	var newSocialmediaId int
	var newSocialmediaCreatedAt, newSocialmediaUpdatedAt string
	err = stmt.QueryRow(socialmedia.UserId, socialmedia.Name, socialmedia.SocialmediaUrl).Scan(
		&newSocialmediaId, &newSocialmediaCreatedAt, &newSocialmediaUpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return err, socialmedia
	}
	socialmedia.SocialmediaId = newSocialmediaId
	socialmedia.CreatedAt = newSocialmediaCreatedAt
	socialmedia.UpdatedAt = newSocialmediaUpdatedAt

	return err, socialmedia
}

func (u socialmediaRepo) getSocialmedia(ctx context.Context, socialmediaId int) (err error, socialmedia Socialmedia, username string) {
	query := `
		SELECT s.socialmedia_id, s.user_id, u.username, s.name, s.socialmedia_url, s.created_at, s.updated_at
		FROM "socialmedia" as s
		JOIN "user" as u
		ON s.user_id = u.user_id 
		WHERE s.deleted = false
		AND u.deleted = false
		AND s.socialmedia_id = $1
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, socialmedia, username
	}
	defer stmt.Close()

	row, err := stmt.QueryContext(ctx, socialmediaId)
	if err != nil {
		log.Println(err)
		return err, socialmedia, username
	}
	if row.Next() {
		err := row.Scan(
			&socialmedia.SocialmediaId,
			&socialmedia.UserId,
			&username,
			&socialmedia.Name,
			&socialmedia.SocialmediaUrl,
			&socialmedia.CreatedAt,
			&socialmedia.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return err, socialmedia, username
		}
	}

	return err, socialmedia, username
}

func (u socialmediaRepo) getSocialmediaList(ctx context.Context) (err error, socialmediaList []Socialmedia) {
	// TODO: select username
	query := `
		SELECT s.socialmedia_id, s.user_id, s.name, s.socialmedia_url, s.created_at, s.updated_at
		FROM "socialmedia" as s
		JOIN "user" as u
		ON s.user_id = u.user_id 
		WHERE s.deleted = false
		AND u.deleted = false
		ORDER BY s.updated_at DESC
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, socialmediaList
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Println(err)
		return err, socialmediaList
	}
	for rows.Next() {
		var socialmedia Socialmedia
		err := rows.Scan(
			&socialmedia.SocialmediaId,
			&socialmedia.UserId,
			&socialmedia.Name,
			&socialmedia.SocialmediaUrl,
			&socialmedia.CreatedAt,
			&socialmedia.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return err, socialmediaList
		}
		socialmediaList = append(socialmediaList, socialmedia)
	}

	return err, socialmediaList
}

func (u socialmediaRepo) editSocialmedia(ctx context.Context, socialmedia Socialmedia) (err error, _ Socialmedia) {
	query := `
		UPDATE "socialmedia" SET 
		name = $1,
		socialmedia_url = $2,
		updated_at = now()
		WHERE socialmedia_id = $3
		AND user_id = $4
		RETURNING created_at, updated_at
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, socialmedia
	}
	defer stmt.Close()

	var newSocialmediaCreatedAt, newSocialmediaUpdatedAt string
	err = stmt.QueryRow(socialmedia.Name, socialmedia.SocialmediaUrl, socialmedia.SocialmediaId, socialmedia.UserId).Scan(
		&newSocialmediaCreatedAt, &newSocialmediaUpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return err, socialmedia
	}
	socialmedia.CreatedAt = newSocialmediaCreatedAt
	socialmedia.UpdatedAt = newSocialmediaUpdatedAt

	return err, socialmedia
}

func (u socialmediaRepo) removeSocialmedia(ctx context.Context, socialmediaId int, userId int) (err error) {
	query := `
		UPDATE "socialmedia" SET updated_at = now(), deleted = true
		WHERE socialmedia_id = $1 AND user_id = $2
	`
	stmt, err := u.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return
	}

	_, err = stmt.ExecContext(ctx, socialmediaId, userId)
	if err != nil {
		log.Println(err)
		return err
	}

	return
}
