package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/kylerequez/go-echo-library/src/database"
	"github.com/kylerequez/go-echo-library/src/repositories"
	"github.com/kylerequez/go-echo-library/src/utils"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("seed failed: %v", err)
	}
}

func run() error {
	if err := utils.LoadEnv(".env.local"); err != nil {
		return err
	}

	if err := database.Connect(); err != nil {
		return err
	}
	defer database.Disconnect()

	ctx := context.Background()

	tx, err := database.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	defer tx.Rollback()

	log.Println("Wiping existing books and authors...")
	if _, err := tx.ExecContext(ctx, "DELETE FROM books"); err != nil {
		return fmt.Errorf("clearing books: %w", err)
	}
	if _, err := tx.ExecContext(ctx, "DELETE FROM authors"); err != nil {
		return fmt.Errorf("clearing authors: %w", err)
	}

	q := repositories.New(tx)

	var authorCount, bookCount, isbnSeq int

	for _, a := range seedAuthors {
		authorID := uuid.New()

		birthDate, err := time.Parse("2006-01-02", a.BirthDate)
		if err != nil {
			return fmt.Errorf("parsing birth date for %s: %w", a.Name, err)
		}

		if err := q.CreateAuthor(ctx, repositories.CreateAuthorParams{
			ID:          authorID,
			Name:        a.Name,
			Biography:   a.Biography,
			BirthDate:   birthDate,
			Nationality: a.Nationality,
			Image:       sql.NullString{},
		}); err != nil {
			return fmt.Errorf("creating author %s: %w", a.Name, err)
		}
		authorCount++

		for _, b := range a.Books {
			isbnSeq++
			isbn := fmt.Sprintf("978000000%05d", isbnSeq)

			publishedDate := time.Date(b.PublishedYear, 1, 1, 0, 0, 0, 0, time.UTC)

			if err := q.CreateBook(ctx, repositories.CreateBookParams{
				ID:            uuid.New(),
				Title:         b.Title,
				AuthorID:      authorID,
				Isbn:          isbn,
				Publisher:     nullString(b.Publisher),
				PublishedDate: publishedDate,
				Pages:         nullInt64(int64(b.Pages)),
				Language:      nullString(b.Language),
				Genre:         nullString(b.Genre),
				Description:   nullString(b.Description),
			}); err != nil {
				return fmt.Errorf("creating book %q: %w", b.Title, err)
			}
			bookCount++
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	log.Printf("Seeded %d authors and %d books.", authorCount, bookCount)
	return nil
}

func nullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func nullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: i > 0}
}
