package user

// import (
// 	"errors"
// 	"testing"

// 	"github.com/alexedwards/argon2id"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type FakeRepository struct {
// 	user User
// 	err  error
// }

// func (f *FakeRepository) Save(user User) (User, error) {
// 	return f.user, f.err
// }

// func (f *FakeRepository) FindByEmail(email string) (User, error) {
// 	return f.user, f.err
// }

// type MockRepository struct {
// 	mock.Mock
// }

// func (m *MockRepository) Save(user User) (User, error) {
// 	args := m.Called(user)
// 	return args.Get(0).(User), args.Error(1)
// }

// func (m *MockRepository) FindByEmail(email string) (User, error) {
// 	args := m.Called(email)
// 	return args.Get(0).(User), args.Error(1)
// }

// func TestRegisterUser(t *testing.T) {
// 	mockRepo := new(MockRepository)
// 	s := NewService(mockRepo)

// 	input := RegisterUserInput{
// 		Name:       "John Doe",
// 		Occupation: "Engineer",
// 		Email:      "john@example.com",
// 		Password:   "password123",
// 	}

// 	mockRepo.On("Save", mock.AnythingOfType("User")).Return(User{ID: 1}, nil)

// 	newUser, err := s.RegisterUser(input)

// 	assert.NoError(t, err)
// 	assert.NotEqual(t, newUser.ID, 0)
// 	mockRepo.AssertExpectations(t)
// }

// func TestRegisterUserError(t *testing.T) {
// 	mockRepo := new(MockRepository)
// 	s := NewService(mockRepo)

// 	input := RegisterUserInput{
// 		Name:       "John Doe",
// 		Occupation: "Engineer",
// 		Email:      "john@example.com",
// 		Password:   "password123",
// 	}

// 	mockRepo.On("Save", mock.AnythingOfType("User")).Return(User{}, errors.New("Failed to save user"))

// 	_, err := s.RegisterUser(input)

// 	assert.Error(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestLogin(t *testing.T) {
// 	mockRepo := new(MockRepository)
// 	s := NewService(mockRepo)

// 	input := LoginInput{
// 		Email:    "john@example.com",
// 		Password: "password123",
// 	}

// 	hashedPassword, _ := argon2id.CreateHash(input.Password, &argon2id.Params{
// 		Memory:      128 * 1024,
// 		Iterations:  4,
// 		Parallelism: 4,
// 		SaltLength:  16,
// 		KeyLength:   32,
// 	})

// 	mockRepo.On("FindByEmail", input.Email).Return(User{
// 		ID:       1,
// 		Email:    input.Email,
// 		Password: hashedPassword,
// 	}, nil)

// 	user, err := s.Login(input)

// 	assert.NoError(t, err)
// 	assert.NotEqual(t, user.ID, 0)
// 	assert.Equal(t, user.Email, input.Email)
// 	mockRepo.AssertExpectations(t)
// }

// func TestLoginErrorUserNotFound(t *testing.T) {
// 	mockRepo := new(MockRepository)
// 	s := NewService(mockRepo)

// 	input := LoginInput{
// 		Email:    "nonexistent@example.com",
// 		Password: "password123",
// 	}

// 	mockRepo.On("FindByEmail", input.Email).Return(User{}, nil)

// 	_, err := s.Login(input)

// 	assert.Error(t, err)
// 	assert.EqualError(t, err, "no user found on that email")
// 	mockRepo.AssertExpectations(t)
// }

// func TestLoginErrorInvalidPassword(t *testing.T) {
// 	mockRepo := new(MockRepository)
// 	s := NewService(mockRepo)

// 	input := LoginInput{
// 		Email:    "john@example.com",
// 		Password: "incorrect_password",
// 	}

// 	hashedPassword, _ := argon2id.CreateHash("correct_password", &argon2id.Params{
// 		Memory:      128 * 1024,
// 		Iterations:  4,
// 		Parallelism: 4,
// 		SaltLength:  16,
// 		KeyLength:   32,
// 	})

// 	mockRepo.On("FindByEmail", input.Email).Return(User{
// 		ID:       1,
// 		Email:    input.Email,
// 		Password: hashedPassword,
// 	}, nil)

// 	_, err := s.Login(input)

// 	assert.Error(t, err)
// 	assert.EqualError(t, err, "password does not match")
// 	mockRepo.AssertExpectations(t)
// }

// func TestLoginErrorRepositoryError(t *testing.T) {
// 	mockRepo := new(MockRepository)
// 	s := NewService(mockRepo)

// 	input := LoginInput{
// 		Email:    "john@example.com",
// 		Password: "password123",
// 	}

// 	mockRepo.On("FindByEmail", input.Email).Return(User{}, errors.New("Failed to fetch user"))

// 	_, err := s.Login(input)

// 	assert.Error(t, err)
// 	assert.EqualError(t, err, "Failed to fetch user")
// 	mockRepo.AssertExpectations(t)
// }

// func TestFormatUser(t *testing.T) {
// 	user := User{
// 		ID:         1,
// 		Name:       "John Doe",
// 		Occupation: "Engineer",
// 		Email:      "john@example.com",
// 	}
// 	token := "some-token"

// 	formatter := FormatUser(user, token)

// 	assert.NotNil(t, formatter)
// 	assert.Equal(t, user.ID, formatter.ID)
// 	assert.Equal(t, user.Name, formatter.Name)
// 	assert.Equal(t, user.Occupation, formatter.Occupation)
// 	assert.Equal(t, user.Email, formatter.Email)
// 	assert.Equal(t, token, formatter.Token)
// }
