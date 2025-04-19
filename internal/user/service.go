package user

import (
	"context"
	"github.com/kmurata08/my-go-chi-oapi-playground/internal/common/errors"
	"github.com/kmurata08/my-go-chi-oapi-playground/internal/gen/user"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ListUsers(ctx context.Context) ([]user.User, error) {
	return []user.User{
		{Id: 1, Name: "Alice", Email: "alice@example.com"},
		{Id: 2, Name: "Bob", Email: "bob@example.com"},
	}, nil
}

func (s *Service) GetUser(ctx context.Context, id int) (*user.User, error) {
	if id == 1 {
		return &user.User{Id: 1, Name: "Alice", Email: "alice@example.com"}, nil
	}
	if id == 2 {
		return &user.User{Id: 2, Name: "Bob", Email: "bob@example.com"}, nil
	}
	return nil, nil
}

func (s *Service) CreateUser(ctx context.Context, req user.CreateUserRequest) (*user.User, error) {
	return &user.User{
		Id:    3,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

func (s *Service) UpdateUser(ctx context.Context, id int, req user.UpdateUserRequest) (*user.User, error) {
	user, err := s.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	user.Name = *req.Name
	user.Email = *req.Email

	return user, nil
}

func (s *Service) DeleteUser(ctx context.Context, id int) error {
	user, err := s.GetUser(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.NewNotFoundError("user_not_found", "User not found")
	}

	return nil
}
