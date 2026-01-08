package service

import (
	"errors"
	"fmt"
	"some_app/model"
	"some_app/service/validation"
	"some_app/store"
	"time"
)

type UserService struct {
	store store.UserRepository
}

func NewUserService(store store.UserRepository) *UserService {
	return &UserService{store: store}
}

func (s *UserService) ChangeProfile(authUserId int, userId int, name string, phone string) (*model.User, error) {
	authUser := s.FindUser(authUserId)
	user := s.FindUser(userId)

	if authUser.IsAdmin || authUser.Id == user.Id {
		if user != nil && !user.IsAdmin {
			user.SetLastViewedAt(time.Now())
		}

		if errs := s.CheckUserParam(user, name, phone); len(errs) > 0 {
			return nil, errors.New("fail user param")
		}
	} else {
		return nil, errors.New("not enough rights")
	}

	s.UserSave(user)
	return user, nil
}

func (s *UserService) FindUser(id int) *model.User {
	return s.store.Find(fmt.Sprintf("SELECT * FROM user WHERE id = %d", id))
}

func (s *UserService) CheckUserParam(user *model.User, name string, phone string) (errs map[string]string) {
	errs = make(map[string]string)

	if err := validation.ValidateName(name); err != nil {
		errs[name] = err.Error()
	} else {
		user.Name = name
	}
	if err := validation.ValidatePhone(phone); err != nil {
		errs[phone] = err.Error()
	} else {
		user.Phone = phone
	}

	return errs
}

func (s *UserService) UserSave(user *model.User) {
	s.store.Save(user)
}
