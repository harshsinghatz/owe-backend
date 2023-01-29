-- name: CreateTransaction :one
INSERT INTO Transaction (reciever_id, sender_id, currency, amount, message, deadline, status,type)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetTranaction :one
SELECT * FROM Transaction WHERE id = $1;

-- name: GetAllDebtTransactions :many
SELECT * FROM Transaction WHERE reciever_id = $1 AND status = "accepted";

-- name: GetAllDebtFromAccountId :many
SELECT * FROM Transaction WHERE reciever_id = $1 AND sender_id=$2 AND status = "accepted";

-- name: GetAllLendTransactions :many
SELECT * FROM Transaction WHERE sender_id = $1 AND status = "accepted";

-- name: GetAllLendFromAccountId :many
SELECT * FROM Transaction WHERE sender_id = $1 AND reciever_id=$2 AND status = "accepted";

-- name: UpdateTransactionStatus :exec
UPDATE Transaction SET status = $2 WHERE id = $1;

-- name: UpdateTransactionType :exec
UPDATE Transaction SET type = $2 WHERE id = $1;

-- name: UpdateTransactionAmount :exec
UPDATE Transaction SET amount = $2 WHERE id = $1;

-- name: UpdateDeadline :exec
UPDATE Transaction SET deadline = $2 WHERE id = $1;

-- name: DeleteTransaction :exec
DELETE FROM Transaction WHERE id = $1;