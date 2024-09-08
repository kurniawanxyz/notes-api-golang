package usecase

import (
	"notes/domain"
)

type NoteUsecase struct {
	repo domain.NoteRepository
}

func NewNoteUsecase(repo domain.NoteRepository) *NoteUsecase {
	return &NoteUsecase{repo}
}

func (uc *NoteUsecase) Get() ([]domain.Note, error) {
	return uc.repo.Get()
}

func (uc *NoteUsecase) Show(id uint) (*domain.Note, error) {
	return uc.repo.Show(id)
}

func (uc *NoteUsecase) Store(data *domain.Note) error {
	return uc.repo.Store(data)
}

func (uc *NoteUsecase) Update(data *domain.Note) error {
	return uc.repo.Update(data)
}

func (uc *NoteUsecase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
