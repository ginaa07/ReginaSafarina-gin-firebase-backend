package services

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ReginaSafarina/gin-firebase-backend/config"
	"github.com/ReginaSafarina/gin-firebase-backend/models"
	"github.com/ReginaSafarina/gin-firebase-backend/repositories"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{userRepo: repositories.NewUserRepository()}
}

func (s *AuthService) VerifyFirebaseToken(firebaseToken string) (string, *models.User, error){
	token, err := config.FirebaseAuth.VerifyIDToken(context.Background(), firebaseToken)
	if err != nil {
		return "", nil, errors.New("firebase token tidak valid atau kadaluarsa")
	}

	emailVerified, _ := token.Claims["email_verified"].(bool)
	if !emailVerified {
		return "", nil, errors.New("EMAIL_NOT_VERIFIED")
	}

	uid := token.UID
	email, _ := token.Claims["email"].(string)
	name, _ := token.Claims["name"].(string)

	user, err := s.userRepo.FindByFirebaseUID(uid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		now := time.Now().Unix()
		user = &models.User{
			FirebaseUID:   uid,
			Email:         email,
			Name:          name,
			Role: 		   "user",
			EmailVerified: true,
			LastLoginAt:   &now,
		}