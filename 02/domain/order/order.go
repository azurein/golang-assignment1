package order

type Order struct {
	OrderId      int    `db:"order_id"`
	CustomerName string `db:"customer_name"`
	OrderedAt    string `db:"ordered_at"`
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
}

type OrderDetail struct {
	OrderDetailId int    `db:"order_detail_id"`
	OrderId       int    `db:"order_id"`
	ItemCode      string `db:"item_code"`
	Description   string `db:"description"`
	Quantity      int    `db:"quantity"`
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
}
