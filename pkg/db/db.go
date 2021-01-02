package db

import (
	"github.com/emilgelman/gocrud/pkg/domain"
)

type Db interface {
	Get(id string) (domain.Article, error)
	GetAll() []domain.Article
	Create(id string, article domain.Article)
	Update(id string, article domain.Article) error
	Delete(id string) error
}
