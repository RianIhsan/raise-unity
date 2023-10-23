package user

import (
	"errors"
	"time"

	"github.com/RianIhsan/raise-unity/helper"
	"github.com/alexedwards/argon2id"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserByEmail(email string) (User, error)        // Tambahkan metode GetUserByEmail
	FindValidOTP(userID int, otp string) (OTP, error) // Tambahkan metode FindValidOTP
	UpdateUser(user User) (User, error)               // Tambahkan metode UpdateUser
	VerifyEmail(email string, otp string) error
	ResendOTP(email string) (OTP, error)
	SaveAvatar(userID int, file string) (User, error)
	GetUserByID(ID int) (User, error)
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
	user.Avatar = "https://res.cloudinary.com/dyominih0/image/upload/v1697817852/default-avatar-icon-of-social-media-user-vector_p8sqa6.jpg"
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, errors.New("error saving user")
	}

	otp := helper.GenerateRandomOTP(6)

	otpModel := OTP{
		UserID:     newUser.ID,
		OTP:        otp,
		ExpiredOTP: time.Now().Add(2 * time.Minute).Unix(),
	}

	_, errOtp := s.repository.SaveOTP(otpModel)
	if errOtp != nil {
		return newUser, errors.New("error sending OTP")
	}

	err = helper.SendOTPByEmail(newUser.Email, otp)
	if err != nil {
		return newUser, errors.New("error sending OTP")
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
	if !user.IsVerified {
		return user, errors.New("account has not been verified")
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

func (s *service) VerifyEmail(email string, otp string) error {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("no user found on that email")
	}

	otpModel, err := s.repository.FindValidOTP(user.ID, otp)
	if err != nil {
		return errors.New("invalid or expired OTP")
	}

	if otpModel.ID == 0 {
		return errors.New("invalid or expired OTP")
	}

	user.IsVerified = true

	_, errUpdate := s.repository.UpdateUser(user)
	if errUpdate != nil {
		return errors.New("error updating user")
	}

	errDeleteOTP := s.repository.DeleteOTP(otpModel)
	if errDeleteOTP != nil {
		return errors.New("error deleting OTP")
	}

	return nil
}

func (s *service) GetUserByEmail(email string) (User, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) FindValidOTP(userID int, otp string) (OTP, error) {
	validOTP, err := s.repository.FindValidOTP(userID, otp)
	if err != nil {
		return validOTP, err
	}
	return validOTP, nil
}

func (s *service) UpdateUser(user User) (User, error) {
	updatedUser, err := s.repository.UpdateUser(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

func (s *service) ResendOTP(email string) (OTP, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return OTP{}, err
	}

	if user.ID == 0 {
		return OTP{}, errors.New("no user found with that email")
	}

	// Hapus OTP sebelumnya
	errDel := s.repository.DeleteUserOTP(user.ID)
	if errDel != nil {
		return OTP{}, errDel
	}

	otp := helper.GenerateRandomOTP(6)
	otpModel := OTP{
		UserID:     user.ID,
		OTP:        otp,
		ExpiredOTP: time.Now().Add(2 * time.Minute).Unix(),
	}

	_, errOtp := s.repository.SaveOTP(otpModel)
	if errOtp != nil {
		return OTP{}, errOtp
	}

	return otpModel, nil
}

func (s *service) SaveAvatar(userID int, file string) (User, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return user, err
	}

	user.Avatar = file

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on with that ID")
	}
	return user, nil
}
