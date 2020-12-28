package models

import "time"

type Post struct {
	ID uint
	Title string 		`json:"title" binding:"required"`
	Body string 		`json:"body"  binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Published bool
	Views uint
}
