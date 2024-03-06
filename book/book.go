package book

type GenreType string

const (
	Terror         GenreType = "terror"
	Romance        GenreType = "romance"
	ScienceFiction GenreType = "science-fiction"
	Novel          GenreType = "novel"
)

type Book struct {
	ID     int
	Name   string
	Genre  GenreType
	Author string
}

type UseCase interface {
	Get(id int) (Book, error)
	GetAll() ([]Book, error)
	Create(title string, author string, pages int, quantity int) (int, error)
}

type Repository interface {
	Reader
	Writer
}

type Reader interface {
	Get(id int) (Book, error)
	GetAll() ([]Book, error)
}

type Writer interface {
	Create(e *Book) (int, error)
}
