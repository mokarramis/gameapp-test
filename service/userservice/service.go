package userservice

import (
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phonenumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name  string
	Phone string
}

type RegisterResponse struct {
	User entity.User
}

func (s Service) Register(r RegisterRequest) (RegisterResponse, error) {
	// TODO: verify phone by verification code

	if !phonenumber.IsValid(r.Phone) {
		return RegisterResponse{}, fmt.Errorf("phone is not valid")
	}
	if isUnique, error := s.repo.IsPhoneNumberUnique(r.Phone); error != nil || !isUnique {
		if error != nil {
			return RegisterResponse{}, fmt.Errorf("unexoected error: %w", error)
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone is not unique")
		}
	}

	// validate name
	if len(r.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name is not valid")
	}

	user := entity.User{
		ID:    0,
		Name:  r.Name,
		Phone: r.Phone,
	}

	createdUser, error := s.repo.Register(user)
	if error != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", error)
	}

	return RegisterResponse{
		User: createdUser,
	}, nil

}
