-- +goose Up
CREATE TABLE authors (
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    biography TEXT NOT NULL,
    birth_date DATETIME NOT NULL,
    nationality TEXT NOT NULL,
    image TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE authors;
