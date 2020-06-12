package services

import (
	"golang.org/x/crypto/bcrypt"
	"mvcApiTest/pkg/models"
	"mvcApiTest/pkg/repositories"
)

type UserService struct {
	repo *repositories.UsersRepository
}

func generateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func NewUserService( repo *repositories.UsersRepository ) *UserService {
	return &UserService{repo:repo}
}

func (s *UserService)GetAll() []models.User {
	userRepository := repositories.NewUsersRepository()
	return userRepository.FetchAll()
}

func (s *UserService)IsUserExists(email string) (uint, error){
	userRepository := repositories.NewUsersRepository()
	count, err := userRepository.UserCount(email)
	if err != nil{
		return 0, err
	}
	return count, nil
}


func (s *UserService)InsertUser(user models.User) error {
	userRepository := repositories.NewUsersRepository()
	hashPassword, err := generateHashPassword(user.Password)
	if err != nil{
		return err
	}
	user.Password = hashPassword
	err = userRepository.Store(user)
	if err != nil{
		return err
	}
	return nil
}