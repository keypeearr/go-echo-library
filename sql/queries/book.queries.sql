-- name: GetBooks :many
SELECT
    *
FROM
    books
LIMIT
    ?
OFFSET
    ?;

-- name: GetBookById :one
SELECT
    *
FROM
    books
WHERE
    id = ?;

-- name: GetBooksCount :one
SELECT
    COUNT(*)
FROM
    books;

-- name: GetBooksCountByAuthor :one
SELECT
    COUNT(*)
FROM
    books
WHERE
    author_id = ?;

-- name: GetBooksByAuthor :many
SELECT
    *
FROM
    books
WHERE
    author_id = ?
LIMIT
    ?
OFFSET
    ?;

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
DELETE FROM books
WHERE
    id = ?;
