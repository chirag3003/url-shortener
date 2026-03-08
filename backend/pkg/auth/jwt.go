package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTService handles JWT token creation and validation.
type JWTService struct {
	secret     []byte
	expiration time.Duration
}

// Claims represents the custom JWT claims.
type Claims struct {
	UserID  string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	PhoneNo string `json:"phoneNo"`
	jwt.RegisteredClaims
}

// NewJWTService creates a new JWTService.
func NewJWTService(secret string, expiration time.Duration) *JWTService {
	return &JWTService{
		secret:     []byte(secret),
		expiration: expiration,
	}
}

// GenerateToken creates a signed JWT for the given user details.
func (s *JWTService) GenerateToken(userID, name, email, phoneNo string) (string, error) {
	now := time.Now()
	claims := &Claims{
		UserID:  userID,
		Name:    name,
		Email:   email,
		PhoneNo: phoneNo,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.expiration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

// ParseToken validates and parses a JWT token string, returning the claims.
func (s *JWTService) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
