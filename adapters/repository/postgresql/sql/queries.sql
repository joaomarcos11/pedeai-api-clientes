-- ----------------------------------------------
-- Clientes

-- name: GetClienteById :one
SELECT * FROM clientes WHERE id = $1 LIMIT 1;

-- name: GetClienteByCPF :one
SELECT * FROM clientes WHERE cpf = $1 LIMIT 1;

-- name: GetClienteByEmail :one
SELECT * FROM clientes WHERE email = $1 LIMIT 1;

-- name: ListCliente :many
SELECT * FROM clientes ORDER BY cpf;

-- name: CreateCliente :one
INSERT INTO  clientes
(id, nome, cpf, email, ativo)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateCliente :exec
UPDATE clientes SET
(nome, cpf, email, ativo) = ($2, $3, $4, $5)
WHERE id = $1;

-- name: DeleteCliente :exec
DELETE FROM clientes WHERE id = $1;

-- name: DeleteAllCliente :exec
DELETE FROM clientes;