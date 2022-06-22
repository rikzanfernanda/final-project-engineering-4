package usecase

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rg-km/final-project-engineering-4/backend/domain"
	"golang.org/x/crypto/bcrypt"
)

type guruUseCase struct {
	guruRepo domain.GuruRepository
}

func NewGuruUseCase(guruRepo domain.GuruRepository) domain.GuruUseCase {
	return &guruUseCase{
		guruRepo: guruRepo,
	}
}

func (g *guruUseCase) Register(guru domain.Guru) (*domain.Guru, error) {
	_, err := g.guruRepo.GetByEmail(guru.Email)
	if err == nil {
		return nil, domain.ErrEmailExists
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(guru.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	guru.Password = string(hashedPass)

	err = g.guruRepo.Create(guru)
	if err != nil {
		return nil, err
	}
	return g.guruRepo.GetByEmail(guru.Email)
}
func (g *guruUseCase) Login(email, password string) (*domain.Guru, string, error) {
	guru, err := g.guruRepo.GetByEmail(email)
	if err != nil {
		return nil, "", domain.ErrEmailNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(guru.Password), ([]byte(password)))
	if err != nil {
		return nil, "", domain.ErrPasswordWrong
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   guru.ID,
		"role": "Guru",
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, "", err
	}

	return guru, tokenString, nil
}
