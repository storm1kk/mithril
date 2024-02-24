package storage

import "github.com/storm1kk/mithril/internal/entity"

type Storage interface {
	CreateUser(user entity.User) (int64, error)
	GetUser(id int64) (entity.User, error)
}
