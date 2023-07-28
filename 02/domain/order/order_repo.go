package order

import (
	"context"
	"database/sql"
	"log"
)

type orderRepo struct {
	db *sql.DB
}

func newOrderRepo(db *sql.DB) orderRepo {
	return orderRepo{
		db: db,
	}
}

func (u orderRepo) createOrder(ctx context.Context, order Order, orderDetailList []OrderDetail) (err error, _ Order, _ []OrderDetail) {
	trx, err := u.db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	query := `
		INSERT INTO "order" (customer_name, ordered_at, created_at, updated_at)
		VALUES ($1, $2, now(), now())
		RETURNING order_id, created_at, updated_at
	`
	stmt, err := trx.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, order, orderDetailList
	}

	var newOrderId int
	var newOrderCreatedAt, newOrderUpdatedAt string
	err = stmt.QueryRow(order.CustomerName, order.OrderedAt).Scan(&newOrderId, &newOrderCreatedAt, &newOrderUpdatedAt)
	if err != nil {
		log.Println(err)
		return err, order, orderDetailList
	}
	order.OrderId = newOrderId
	order.CreatedAt = newOrderCreatedAt
	order.UpdatedAt = newOrderUpdatedAt

	for idx, orderDetail := range orderDetailList {
		query = `
			INSERT INTO "order_detail" (order_id, item_code, description, quantity, created_at, updated_at)
			VALUES ($1, $2, $3, $4, now(), now())
			RETURNING order_detail_id, created_at, updated_at
		`
		stmt, err = trx.Prepare(query)
		if err != nil {
			log.Println(err)
			return err, order, orderDetailList

		}
		var newOrderDetailId int
		var newOrderCreatedAt, newOrderUpdatedAt string
		err = stmt.QueryRow(newOrderId, orderDetail.ItemCode, orderDetail.Description, orderDetail.Quantity).Scan(&newOrderDetailId, &newOrderCreatedAt, &newOrderUpdatedAt)
		if err != nil {
			log.Println(err)
			return err, order, orderDetailList
		}
		orderDetailList[idx].OrderDetailId = newOrderDetailId
		orderDetailList[idx].CreatedAt = newOrderCreatedAt
		orderDetailList[idx].UpdatedAt = newOrderUpdatedAt
	}

	err = trx.Commit()
	if err != nil {
		log.Println(err)
		return err, order, orderDetailList
	}
	stmt.Close()

	return err, order, orderDetailList
}

func (u orderRepo) getOrderList(ctx context.Context) (err error, orderList []Order, orderDetailList []OrderDetail) {
	query := `
		SELECT order_id, customer_name, ordered_at, created_at, updated_at
		FROM "order"
	`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, orderList, orderDetailList
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Println(err)
		return err, orderList, orderDetailList
	}
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.OrderId,
			&order.CustomerName,
			&order.OrderedAt,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return err, orderList, orderDetailList
		}
		orderList = append(orderList, order)
	}

	query = `
		SELECT order_detail_id, order_id, item_code, description, quantity, created_at, updated_at
		FROM "order_detail"
	`
	stmt, err = u.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, orderList, orderDetailList
	}
	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		log.Println(err)
		return err, orderList, orderDetailList
	}
	for rows.Next() {
		var orderDetail OrderDetail
		err := rows.Scan(
			&orderDetail.OrderDetailId,
			&orderDetail.OrderId,
			&orderDetail.ItemCode,
			&orderDetail.Description,
			&orderDetail.Quantity,
			&orderDetail.CreatedAt,
			&orderDetail.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return err, orderList, orderDetailList
		}
		orderDetailList = append(orderDetailList, orderDetail)
	}

	stmt.Close()
	return err, orderList, orderDetailList
}

func (u orderRepo) editOrder(ctx context.Context, order Order, orderDetailList []OrderDetail) (err error, _ Order, _ []OrderDetail) {
	trx, err := u.db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	query := `
		UPDATE "order" SET 
			customer_name = $1, 
			ordered_at = $2, 
			updated_at = now()
		WHERE order_id = $3
		RETURNING created_at, updated_at
	`
	stmt, err := trx.Prepare(query)
	if err != nil {
		log.Println(err)
		return err, order, orderDetailList
	}
	var newOrderCreatedAt, newOrderUpdatedAt string
	err = stmt.QueryRow(order.CustomerName, order.OrderedAt, order.OrderId).Scan(&newOrderCreatedAt, &newOrderUpdatedAt)
	if err != nil {
		log.Println(err)
		return err, order, orderDetailList
	}
	order.CreatedAt = newOrderCreatedAt
	order.UpdatedAt = newOrderUpdatedAt

	for idx, orderDetail := range orderDetailList {
		query = `
			UPDATE "order_detail" SET
				item_code = $1, 
				description = $2, 
				quantity = $3, 
				updated_at = now()
			WHERE order_detail_id = $4
			RETURNING created_at, updated_at
		`
		stmt, err = trx.Prepare(query)
		if err != nil {
			log.Println(err)
			return err, order, orderDetailList

		}
		var newOrderCreatedAt, newOrderUpdatedAt string
		err = stmt.QueryRow(orderDetail.ItemCode, orderDetail.Description, orderDetail.Quantity, orderDetail.OrderDetailId).Scan(&newOrderCreatedAt, &newOrderUpdatedAt)
		if err != nil {
			log.Println(err)
			return err, order, orderDetailList
		}
		orderDetailList[idx].CreatedAt = newOrderCreatedAt
		orderDetailList[idx].UpdatedAt = newOrderUpdatedAt
	}

	err = trx.Commit()
	if err != nil {
		log.Println(err)
		return err, order, orderDetailList
	}
	stmt.Close()

	return err, order, orderDetailList
}

func (u orderRepo) removeOrder(ctx context.Context, order Order) (err error) {
	trx, err := u.db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	query := `
		DELETE FROM "order"
		WHERE order_id = $1
	`
	stmt, err := trx.Prepare(query)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = stmt.ExecContext(ctx, order.OrderId)
	if err != nil {
		log.Println(err)
		return err
	}

	query = `
		DELETE FROM "order_detail"
		WHERE order_id = $1
	`
	stmt, err = trx.Prepare(query)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = stmt.ExecContext(ctx, order.OrderId)
	if err != nil {
		log.Println(err)
		return err
	}

	err = trx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	stmt.Close()

	return
}
