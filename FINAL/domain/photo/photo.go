package photo

type Photo struct {
	PhotoId   int    `db:"photo_id"`
	UserId    int    `db:"user_id"`
	Title     string `db:"title"`
	Caption   string `db:"caption"`
	PhotoUrl  string `db:"photo_url"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	Deleted   bool   `db:"deleted"`
}
