-- name: CreateAccount :one
INSERT INTO Accounts (
  name, phone_number,currency,balance
) VALUES (
  $1, $2, $3, $4
) RETURNING *;
 
-- name: GetAccount :one
SELECT * FROM Accounts 
WHERE id = $1 LIMIT 1;

-- name: UpdateAccountPhoneNumber :exec
UPDATE Accounts
SET phone_number = $2
WHERE id = $1;

-- name: UpdateAccountBalance :exec
UPDATE Accounts
SET balance = $2
WHERE id = $1;