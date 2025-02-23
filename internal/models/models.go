package models

import (
	"time"
)

type Author struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" binding:"required"`
	Biography string    `json:"biography"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
	Books     []Book    `json:"books,omitempty"`
}

type Book struct {
	ID              uint     `gorm:"primaryKey" json:"id"`
	Title           string   `json:"title" binding:"required"`
	AuthorID        uint     `json:"author_id" binding:"required"`
	ISBN            string   `json:"isbn" binding:"required"`
	PublicationYear int      `json:"publication_year" binding:"required"`
	Description     string   `json:"description"`
	Author          Author   `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Reviews         []Review `gorm:"foreignKey:BookID" json:"reviews,omitempty"`
}

type Review struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	BookID     uint      `json:"book_id" binding:"required"`
	Rating     int       `json:"rating" binding:"required,min=1,max=5"`
	Comment    string    `json:"comment" binding:"required"`
	DatePosted time.Time `json:"date_posted" binding:"required"`
}
