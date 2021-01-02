package db

import (
	"errors"
	"github.com/emilgelman/gocrud/pkg/domain"
)


type Memory struct {
	data map[string]domain.Article
}
const NoSuchKeyError string = "no such key"

func NewMemory() *Memory {
	return &Memory{data: make(map[string]domain.Article)}
}

func (m Memory) Get(id string) (domain.Article, error) {
	if value, res := m.data[id]; res {
		return value, nil
	}
	return domain.Article{}, errors.New(NoSuchKeyError)
}

func (m Memory) GetAll() []domain.Article {
	var res []domain.Article
	for _, v := range m.data {
		res = append(res, v)
	}
	return res
}

func (m Memory) Create(id string, article domain.Article) {
	m.data[id] = article
}

func (m Memory) Update(id string, article domain.Article) error {
	if _, exists := m.data[id]; !exists {
		return errors.New(NoSuchKeyError)
	}
	m.data[id] = article
	return nil
}

func (m Memory) Delete(id string) error {
	if _, exists:= m.data[id]; !exists {
		return errors.New(NoSuchKeyError)
	}
	delete(m.data, id)
	return nil
}
