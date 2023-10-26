package user

//import (
//	"github.com/RianIhsan/raise-unity/user/mocks"
//	"github.com/alexedwards/argon2id"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"testing"
//)
//
//func TestRegisterUser(t *testing.T) {
//	repoMock := mocks.NewRepoAuthInterface(t)
//	service := NewService(repoMock)
//
//	repoMock.On("Save", mock.Anything).Return(User{}, nil)
//	repoMock.On("SaveOTP", mock.Anything).Return(OTP{}, nil)
//
//	input := RegisterUserInput{
//		Name:       "ExampleTest",
//		Occupation: "Programmer",
//		Email:      "Example@gmail.com",
//		Password:   "ExampleTest",
//	}
//
//	user, err := service.RegisterUser(input)
//	assert.NotNil(t, user)
//	assert.Error(t, err)
//
//	repoMock.AssertExpectations(t)
//}
//
//func TestLogin(t *testing.T) {
//	repoMock := mocks.NewRepoAuthInterface(t)
//	service := NewService(repoMock)
//
//	expectedEmail := "example@gmail.com"
//	passwordHash, _ := argon2id.CreateHash("RianTest123", &argon2id.Params{
//		Memory:      128 * 1024,
//		Iterations:  4,
//		Parallelism: 4,
//		SaltLength:  16,
//		KeyLength:   32,
//	})
//	expectedUser := User{
//		ID:         1,
//		Email:      expectedEmail,
//		Password:   passwordHash,
//		IsVerified: true,
//	}
//
//	repoMock.On("FindByEmail", "example@gmail.com").Return(expectedUser, nil)
//
//	//correctPassword,_ := argon2id.ComparePasswordAndHash(passwordHash, expectedUser.Password)
//	incorrectPassword := "wrongpassword"
//
//	user, err := service.Login(LoginInput{
//		Email:    expectedEmail,
//		Password: expectedUser.Password,
//	})
//
//	assert.NotNil(t, user)
//	assert.NoError(t, err)
//
//	user, err = service.Login(LoginInput{
//		Email:    expectedEmail,
//		Password: incorrectPassword,
//	})
//
//	assert.NotNil(t, user)
//	assert.Nil(t, err)
//	assert.Nil(t, err)
//
//	repoMock.AssertExpectations(t)
//
//}
