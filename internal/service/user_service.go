package service

import (
	"database/sql"
	"errors"
	"net"
	"sj/internal/db/sqlc"
	"sj/internal/dto"
	"sj/internal/repository"
	"sj/pkg/bcrypt"
	"sj/pkg/jwt"
	"strings"

	"github.com/google/uuid"
)

type UserService interface {
	AddUser(arg dto.UserCreateReq) (sql.Result, error)
	EmailExists(email string) (bool, error)
	Login(req dto.UserLoginReq) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) AddUser(arg dto.UserCreateReq) (sql.Result, error) {
	isExist, err := s.EmailExists(arg.Email)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, errors.New("email sudah digunakan")
	}

	domain := strings.Split(arg.Email, "@")[1]
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return nil, err
	}

	hashedPassword, err := bcrypt.HashValue(arg.Password)
	if err != nil {
		return nil, err
	}

	input := sqlc.AddUserParams{
		ID:       uuid.New().String(),
		Email:    arg.Email,
		Password: sql.NullString{String: hashedPassword},
		OauthID:  sql.NullString{},
		Name:     arg.Name,
		Role:     "",
	}
	return s.repo.AddUser(input)
}

func (s *userService) EmailExists(email string) (bool, error) {
	count, err := s.repo.EmailExists(email)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *userService) Login(req dto.UserLoginReq) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.ValidateHash(req.Password, user.Password.String)
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
