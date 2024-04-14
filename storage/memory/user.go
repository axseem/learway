package memory

import (
	"context"
	"time"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/storage"
)

type UserStorage []model.User

func (s UserStorage) GetByID(ctx context.Context, id string) (model.User, error) {
	for _, user := range s {
		if user.ID == id {
			return user, nil
		}
	}
	return model.User{}, ErrNotFound
}

func (s UserStorage) GetByEmail(ctx context.Context, email string) (model.User, error) {
	for _, user := range s {
		if user.Email == email {
			return user, nil
		}
	}
	return model.User{}, ErrNotFound
}

func (s UserStorage) GetByUsername(ctx context.Context, username string) (model.User, error) {
	for _, user := range s {
		if user.Username == username {
			return user, nil
		}
	}
	return model.User{}, ErrNotFound
}

func (s *UserStorage) Create(ctx context.Context, arg storage.UserCreateParams) error {
	newUser := model.User{
		ID:        arg.ID,
		Username:  arg.Username,
		Email:     arg.Email,
		Password:  arg.Password,
		CreatedAt: time.Now(),
	}
	*s = append(*s, newUser)
	return nil
}

func (s *UserStorage) UpdatePassword(ctx context.Context, id string, password []byte) error {
	for i, user := range *s {
		if user.ID == id {
			(*s)[i].Password = password
			return nil
		}
	}
	return ErrNotFound
}

func (s *UserStorage) UpdateUsername(ctx context.Context, id, username string) error {
	for i, user := range *s {
		if user.ID == id {
			(*s)[i].Username = username
			return nil
		}
	}
	return ErrNotFound
}

func (s *UserStorage) UpdateProfile(ctx context.Context, id string, arg model.UserUpdateProfileParams) error {
	for i, user := range *s {
		if user.ID == id {
			(*s)[i].Name = arg.Name
			(*s)[i].Description = arg.Description
			(*s)[i].Picture = arg.Picture
			return nil
		}
	}
	return ErrNotFound
}

func (s *UserStorage) Delete(ctx context.Context, id string) error {
	for i, user := range *s {
		if user.ID == id {
			*s = append((*s)[:i], (*s)[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
