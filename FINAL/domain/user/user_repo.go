package user

import (
	"context"
	"database/sql"
	"log"
)

type userRepo struct {
	db *sql.DB
}

func newUserRepo(db *sql.DB) userRepo {
	return userRepo{
		db: db,
	}
}

func (u userRepo) createUser(ctx context.Context, user User) (err error, _ User) {
	query := `
		INSERT INTO "user" (username, email, password, age, created_at, updated_at)
		VALUES ($1, $2, $3, $4, now(), now())
		RETURNING user_id, created_at, updated_at
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, user
	}
	defer stmt.Close()

	var newUserId int
	var newUserCreatedAt, newUserUpdatedAt string
	err = stmt.QueryRow(user.Username, user.Email, user.Password, user.Age).Scan(
		&newUserId, &newUserCreatedAt, &newUserUpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return err, user
	}
	user.UserId = newUserId
	user.CreatedAt = newUserCreatedAt
	user.UpdatedAt = newUserUpdatedAt

	return err, user
}

func (u userRepo) getUser(ctx context.Context, username string, password string) (err error, user User) {
	query := `
		SELECT user_id, username, password, email, age, created_at, updated_at
		FROM "user"
		WHERE deleted = false
		AND username = $1
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, user
	}
	defer stmt.Close()

	row, err := stmt.QueryContext(ctx, username)
	if err != nil {
		log.Println(err)
		return err, user
	}
	if row.Next() {
		err := row.Scan(
			&user.UserId,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Age,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return err, user
		}
	}

	return err, user
}
