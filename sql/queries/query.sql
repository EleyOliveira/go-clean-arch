-- name: ListOrders :many
SELECT * FROM orders;

-- name: CreateOrder :exec
INSERT INTO orders (ID, Price, Tax, FinalPrice) VALUES (?,?,?,?)