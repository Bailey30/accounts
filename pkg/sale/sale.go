package sale

import (
	"fmt"
	"time"
)

type Sale struct {
	Id          int       `json:"id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Sales struct {
	Sales []Sale `'json:"sales"`
}

type Repo interface {
	Get(id int) (Sale, error)
	GetAll() ([]Sale, error)
	Create(sale Sale) (string, error)
	Update(sale Sale) error
	Delete(id string) error
}

type Service interface {
	Get(id int) (Sale, error)
	GetAll() ([]Sale, error)
	Create(sale Sale) (string, error)
	Update(sale Sale) error
	Delete(id string) error
}

type sale struct {
	repo Repo
}

func New(repo Repo) Service {
	return &sale{repo}
}

func (s *sale) Get(id int) (Sale, error) {
	return s.repo.Get(id)
}

func (s *sale) GetAll() ([]Sale, error) {
	sales, err := s.repo.GetAll()
	if err != nil {
		fmt.Printf("Error getting sale details: %v\n", err)
		return sales, err
	}
	return sales, nil
}

func (s *sale) Create(sale Sale) (string, error) {
	return s.repo.Create(sale)
}

func (s *sale) Update(sale Sale) error {
	return s.repo.Update(sale)
}

func (s *sale) Delete(id string) error {
	return s.repo.Delete(id)
}
