-- name: GetAuthors :many
select *
from authors
limit ?
offset ?
;

-- name: CreateAuthor :exec
INSERT INTO authors (
    id,
    name,
    biography,
    birth_date,
    nationality,
    image
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
);

-- name: UpdateAuthor :one
UPDATE authors SET 
    name = COALESCE(sqlc.narg('name'), name),
    biography = COALESCE(sqlc.narg('biography'), biography),
    birth_date = COALESCE(sqlc.narg('birth_date'), birth_date),
    nationality = COALESCE(sqlc.narg('nationality'), nationality),
    image = COALESCE(sqlc.narg('image'), image)
WHERE
    id = sqlc.arg('id')
RETURNING
    *;

-- name: GetAuthorsCount :one
select count(*)
from authors
;

-- name: GetAuthorByName :one
select *
from authors
where name = ?
;

-- name: GetAuthorById :one
select *
from authors
where id = ?
;

-- name: DeleteAuthorById :exec
delete from authors
where id = ?
;
