package user

import (
	"errors"

	"github.com/alexedwards/argon2id"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := argon2id.CreateHash(input.Password, &argon2id.Params{
		Memory:      128 * 1024,
		Iterations:  4,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	})
	if err != nil {
		return user, errors.New("error creating password hash")
	}
	user.Password = passwordHash
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, errors.New("error saving user")
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}
	match, err := argon2id.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		return user, err
	}
	if !match {
		return user, errors.New("password does not match")
	}
	return user, nil
}
