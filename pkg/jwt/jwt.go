package jwt

import (
	"time"

	"github.com/ak-repo/stream-hub/pkg/errors"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	secretKey       []byte
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewJWTManager(secret string, accessExp, refreshExp time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:       []byte(secret),
		accessTokenTTL:  accessExp,
		refreshTokenTTL: refreshExp,
	}
}

// GenerateAccessToken creates a new JWT access token
func (m *JWTManager) GenerateAccessToken(userID, email string) (string, time.Time, error) {
	expirationTime := time.Now().Add(m.accessTokenTTL)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

// ValidateToken validates and parses a JWT token
func (m *JWTManager) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrInvalidToken
		}
		return m.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.ErrInvalidToken
	}

	// Check expiration
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.ErrExpiredToken
	}

	return claims, nil
}

// GenerateRefreshToken creates a refresh token (simpler, longer-lived)
func (m *JWTManager) GenerateRefreshToken(userID string) (string, time.Time, error) {
	expirationTime := time.Now().Add(m.refreshTokenTTL)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}
