package security

import (
	"pos-backend/internal/domain"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type TokenService interface {
	CreateAuthToken(userID, typeUser string) (*domain.Token, error)
}
type jwtService struct{}

func NewJWTService() TokenService {
	return &jwtService{}
}
func (j *jwtService) CreateAuthToken(id, typeUser string) (*domain.Token, error) {

	refId := uuid.New().String()
	env := viper.GetString("env")

	sessionTimeAccess := 1440   // 1 วัน 1440
	sessionTimeRefresh := 43200 // 3 เดือน 129600

	// ===== Access Token =====
	token := jwt.New(jwt.SigningMethodHS256)
	expAccess := time.Now().Add(time.Minute * time.Duration(sessionTimeAccess)).Unix()
	claims := jwt.MapClaims{
		"type":      "access",
		"env":       env,
		"id":        id,
		"type_user": typeUser,
		"ref_id":    refId,
		"exp":       expAccess,
	}

	token.Claims = claims

	accessToken, err := token.SignedString([]byte(viper.GetString("auth.access")))
	if err != nil {
		return nil, err
	}

	// ===== Refresh Token =====
	tokenRefresh := jwt.New(jwt.SigningMethodHS256)
	expRefesh := time.Now().Add(time.Minute * time.Duration(sessionTimeRefresh)).Unix()
	claimsRefresh := jwt.MapClaims{
		"type":      "refresh",
		"env":       env,
		"id":        id,
		"type_user": typeUser,
		"ref_id":    refId,
		"exp":       expRefesh,
	}

	tokenRefresh.Claims = claimsRefresh

	refreshToken, err := tokenRefresh.SignedString([]byte(viper.GetString("auth.refresh")))
	if err != nil {
		return nil, err
	}

	return &domain.Token{
		Id:         id,
		Access:     accessToken,
		Refresh:    refreshToken,
		AccessTTL:  time.Minute * time.Duration(sessionTimeAccess),
		RefreshTTL: time.Minute * time.Duration(sessionTimeRefresh),
	}, nil
}
