-- name: ListProdcuts :many
SELECT 
 *
FROM
 products;


-- name: FindProductByid :one
SELECT 
 *
FROM
 products 
WHERE id = $1;
 