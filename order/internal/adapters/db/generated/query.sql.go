// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package generated

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOrder = `-- name: CreateOrder :exec
INSERT INTO "orders" ("total_price", "user_id", "updated_at")
VALUES ($1, $2, $3)
`

type CreateOrderParams struct {
	TotalPrice float64          `json:"total_price"`
	UserID     int32            `json:"user_id"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) error {
	_, err := q.db.Exec(ctx, createOrder, arg.TotalPrice, arg.UserID, arg.UpdatedAt)
	return err
}

const createOrderItems = `-- name: CreateOrderItems :exec
INSERT INTO "order_items" ("product_id", "order_id", "amount", "price")
VALUES ($1, $2, $3, $4)
`

type CreateOrderItemsParams struct {
	ProductID pgtype.Int4   `json:"product_id"`
	OrderID   pgtype.Int4   `json:"order_id"`
	Amount    pgtype.Int4   `json:"amount"`
	Price     pgtype.Float8 `json:"price"`
}

func (q *Queries) CreateOrderItems(ctx context.Context, arg CreateOrderItemsParams) error {
	_, err := q.db.Exec(ctx, createOrderItems,
		arg.ProductID,
		arg.OrderID,
		arg.Amount,
		arg.Price,
	)
	return err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM "orders" WHERE "id" = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteOrder, id)
	return err
}

const deleteOrderItems = `-- name: DeleteOrderItems :exec
DELETE FROM "order_items" WHERE "id" = $1
`

func (q *Queries) DeleteOrderItems(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteOrderItems, id)
	return err
}

const getOrder = `-- name: GetOrder :many
SELECT orders.id, user_id, total_price, created_at, updated_at, order_items.id, product_id, order_id, amount, price FROM "orders" 
LEFT JOIN "order_items" on "orders"."id" = "order_items"."order_id"
WHERE "orders"."id" = $1 ORDER BY "order_items"."id"
`

type GetOrderRow struct {
	ID         int32            `json:"id"`
	UserID     int32            `json:"user_id"`
	TotalPrice float64          `json:"total_price"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
	ID_2       pgtype.Int4      `json:"id_2"`
	ProductID  pgtype.Int4      `json:"product_id"`
	OrderID    pgtype.Int4      `json:"order_id"`
	Amount     pgtype.Int4      `json:"amount"`
	Price      pgtype.Float8    `json:"price"`
}

func (q *Queries) GetOrder(ctx context.Context, id int32) ([]GetOrderRow, error) {
	rows, err := q.db.Query(ctx, getOrder, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrderRow{}
	for rows.Next() {
		var i GetOrderRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TotalPrice,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.ProductID,
			&i.OrderID,
			&i.Amount,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderItems = `-- name: GetOrderItems :many
SELECT id, product_id, order_id, amount, price FROM "order_items" WHERE "id" = $1
`

func (q *Queries) GetOrderItems(ctx context.Context, id int32) ([]OrderItem, error) {
	rows, err := q.db.Query(ctx, getOrderItems, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderItem{}
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.OrderID,
			&i.Amount,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrders = `-- name: GetOrders :many
SELECT orders.id, user_id, total_price, created_at, updated_at, order_items.id, product_id, order_id, amount, price FROM "orders" 
JOIN "order_items" on "orders"."id" = "order_items"."order_id" ORDER BY "orders"."id" LIMIT $1 OFFSET $2
`

type GetOrdersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetOrdersRow struct {
	ID         int32            `json:"id"`
	UserID     int32            `json:"user_id"`
	TotalPrice float64          `json:"total_price"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
	ID_2       int32            `json:"id_2"`
	ProductID  pgtype.Int4      `json:"product_id"`
	OrderID    pgtype.Int4      `json:"order_id"`
	Amount     pgtype.Int4      `json:"amount"`
	Price      pgtype.Float8    `json:"price"`
}

func (q *Queries) GetOrders(ctx context.Context, arg GetOrdersParams) ([]GetOrdersRow, error) {
	rows, err := q.db.Query(ctx, getOrders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrdersRow{}
	for rows.Next() {
		var i GetOrdersRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TotalPrice,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.ProductID,
			&i.OrderID,
			&i.Amount,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrder = `-- name: UpdateOrder :exec
UPDATE "orders" 
SET "total_price" = $1, "user_id" = $2, "created_at" = $3, "updated_at" = $4
WHERE "id" = $5
`

type UpdateOrderParams struct {
	TotalPrice float64          `json:"total_price"`
	UserID     int32            `json:"user_id"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
	ID         int32            `json:"id"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) error {
	_, err := q.db.Exec(ctx, updateOrder,
		arg.TotalPrice,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateOrderItems = `-- name: UpdateOrderItems :exec
UPDATE "order_items" 
SET "product_id" = $1, "amount" = $2, "price" = $3
WHERE "id" = $4
`

type UpdateOrderItemsParams struct {
	ProductID pgtype.Int4   `json:"product_id"`
	Amount    pgtype.Int4   `json:"amount"`
	Price     pgtype.Float8 `json:"price"`
	ID        int32         `json:"id"`
}

func (q *Queries) UpdateOrderItems(ctx context.Context, arg UpdateOrderItemsParams) error {
	_, err := q.db.Exec(ctx, updateOrderItems,
		arg.ProductID,
		arg.Amount,
		arg.Price,
		arg.ID,
	)
	return err
}
