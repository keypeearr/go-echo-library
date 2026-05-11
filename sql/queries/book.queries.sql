-- name: GetBooks :many
select *
from books
limit ?
offset ?
;

-- name: GetBookById :one
select *
from books
where id = ?
;

-- name: GetBooksCount :one
select count(*)
from books
;

-- name: GetBookByIsbn :one
select *
from books
where isbn = ?
;

-- name: GetBooksCountByAuthor :one
select count(*)
from books
where author_id = ?
;

-- name: GetBooksByAuthor :many
select *
from books
where author_id = ?
limit ?
offset ?
;

-- name: CreateBook :exec
INSERT INTO books (
    id,
    title,
    author_id,
    isbn,
    publisher,
    published_date,
    pages,
    language,
    genre,
    description
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
); 

-- name: UpdateBook :one
UPDATE books SET 
    title = COALESCE(sqlc.narg('title'), title),
    isbn = COALESCE(sqlc.narg('isbn'), isbn),
    publisher = COALESCE(sqlc.narg('publisher'), publisher),
    published_date = COALESCE(sqlc.narg('published_date'), published_date),
    pages = COALESCE(sqlc.narg('pages'), pages),
    language = COALESCE(sqlc.narg('language'), language),
    genre = COALESCE(sqlc.narg('genre'), genre),
    description = COALESCE(sqlc.narg('description'), description)
WHERE
    id = sqlc.arg('id')
RETURNING
    *;

-- name: DeleteBook :exec
delete from books
where id = ?
;
