package book

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(name, author, genre string) (int, error) {
	b := &Book{
		Name:   name,
		Author: author,
		Genre:  GenreType(genre),
	}

	res, err := s.repo.Create(b)

	if err != nil {
		return 0, err
	}

	return res, nil
}

func (s *Service) Get(id int) (*Book, error) {
	b, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return &b, err
}

func (s *Service) GetAll() (*[]Book, error) {
	books, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return &books, err
}
