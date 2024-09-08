package domain

type Note struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteRepository interface {
	Get() ([]Note, error)
	Show(id uint) (*Note, error)
	Store(data *Note) error
	Update(data *Note) error
	Delete(id uint) error
}
