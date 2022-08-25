package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Buku, error)
	FindByID(ID int) (Buku, error)
	Create(book Buku) (Buku, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Buku, error) {
	var books []Buku

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindByID(ID int) (Buku, error) {
	var book Buku
	err := r.db.Find(&book, ID).Error
	return book, err
}

func (r *repository) Create(book Buku) (Buku, error) {
	err := r.db.Create(&book).Error

	return book, err
}