package socialmedia

type Socialmedia struct {
	SocialmediaId  int    `db:"socialmedia_id"`
	UserId         int    `db:"user_id"`
	Name           string `db:"name"`
	SocialmediaUrl string `db:"socialmedia_url"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
	Deleted        bool   `db:"deleted"`
}
