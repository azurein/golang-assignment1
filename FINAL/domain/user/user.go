package user

type User struct {
	UserId    int    `db:"userId"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Age       int    `db:"age"`
	CreatedAt string `db:"createdAt"`
	UpdatedAt string `db:"updatedAt"`
	Deleted   bool   `db:"deleted"`
}
