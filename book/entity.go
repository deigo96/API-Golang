package book

import "time"

type Buku struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt		time.Time
}