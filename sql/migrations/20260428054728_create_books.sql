-- +goose Up
CREATE TABLE books (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    author_id UUID NOT NULL,
    isbn TEXT UNIQUE NOT NULL,
    publisher TEXT,
    published_date DATETIME NOT NULL,
    pages INT,
    language TEXT,
    genre TEXT,
    description TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES authors (id)
);

-- +goose Down
DROP TABLE books;
