package usecase

import "notes/domain"

type NoteUsecase struct {
	repo domain.NoteRepository
}

func handleStore(repo domain.NoteRepository) *NoteUsecase {
	return &NoteUsecase{repo}
}

func (uc *NoteUsecase) handleGet() ([]domain.Note, error) {
	return uc.repo.Get()
}

func (uc *NoteUsecase) handleShow(id uint) (*domain.Note, error) {
	return uc.repo.Show(id)
}

func (uc *NoteUsecase) handleStore(data *domain.Note) error {
	return uc.repo.Store(data)
}

func (uc *NoteUsecase) handleUpdate(data *domain.Note) error {
	return uc.repo.Update(data)
}

func (uc *NoteUsecase) handleDelete(id uint) error {
	return uc.repo.Delete(id)
}
