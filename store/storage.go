package store

import (
	"some_app/model"
)

type UserRepository interface {
	Find(sql string) *model.User // как будто, мы тут делаем абстрактное хранилище, но внутри интерфейса сразу завязываемся на sql
	Save(user *model.User)
}
