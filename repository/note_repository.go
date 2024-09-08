package repository

import (
	"gorm.io/gorm"

	"notes/domain"
)

type MySQLNoteRepository struct {
	DB *gorm.DB
}

func NewMySQLNoteRepository(db *gorm.DB) *MySQLNoteRepository {
	return &MySQLNoteRepository{db}
}

func (r *MySQLNoteRepository) Get() ([]domain.Note, error) {
	var notes []domain.Note
	if err := r.DB.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *MySQLNoteRepository) Show(id uint) (*domain.Note, error) {
	var note domain.Note
	if err := r.DB.First(&note, id).Error; err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *MySQLNoteRepository) Store(note *domain.Note) error {
	return r.DB.Create(note).Error
}

func (r *MySQLNoteRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Note{}, id).Error
}

func (r *MySQLNoteRepository) Update(note *domain.Note) error {
	return r.DB.Save(note).Error
}
