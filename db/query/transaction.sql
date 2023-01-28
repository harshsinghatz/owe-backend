-- name: CreateTransaction :one
INSERT INTO Transaction (reciever_id, sender_id, currency, amount, message, deadline)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GiveAmountTransactions :many
SELECT * FROM Transaction WHERE reciever_id = $1 LIMIT 1;

-- name: GetAmountTransactions :many
SELECT * FROM Transaction WHERE id = $1 LIMIT 1;

-- name: UpdateTranactionAmount :exec
UPDATE Transaction SET amount = $2 WHERE id = $1;

-- name: UpdateDeadline :exec
UPDATE Transaction SET deadline = $2 WHERE id = $1;

-- name: DeleteTransaction :exec
DELETE FROM Transaction WHERE id = $1;