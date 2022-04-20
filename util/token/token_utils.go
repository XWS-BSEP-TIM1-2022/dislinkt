package token

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtManager struct {
	secret    string
	issuer    string
	expiresIn time.Duration
}

type Claims struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func NewJwtManager(secret string, expiresIn time.Duration) *JwtManager {
	return &JwtManager{secret: secret, expiresIn: expiresIn}
}

func NewJwtManagerDislinkt(expiresIn time.Duration) *JwtManager {
	return &JwtManager{secret: "dislinkt_secret123", expiresIn: expiresIn}
}

func (manager *JwtManager) GenerateJWT(userId string, email string, role string) (string, error) {
	claims := Claims{
		UserId: userId,
		Email:  email,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    manager.issuer,
			ExpiresAt: time.Now().Add(manager.expiresIn).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(manager.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (manager *JwtManager) GetUserIdFromToken(jwtString string) (string, error) {
	token, err := jwt.ParseWithClaims(jwtString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secret), nil
	})
	if err != nil {
		return "", err
	}
	return token.Claims.(*Claims).UserId, nil
}

func (manager *JwtManager) GetEmailFromToken(jwtString string) (string, error) {
	token, err := jwt.ParseWithClaims(jwtString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secret), nil
	})
	if err != nil {
		return "", err
	}
	return token.Claims.(*Claims).Email, nil
}

func (manager *JwtManager) GetRoleFromToken(jwtString string) (string, error) {
	token, err := jwt.ParseWithClaims(jwtString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secret), nil
	})
	if err != nil {
		return "", err
	}
	return token.Claims.(*Claims).Role, nil
}

func (manager *JwtManager) GetExpirationDateFromToken(jwtString string) (time.Time, error) {
	token, err := jwt.ParseWithClaims(jwtString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secret), nil
	})
	if err != nil {
		return time.Now(), err
	}
	return time.Unix(token.Claims.(*Claims).ExpiresAt, 0), nil
}

func (manager *JwtManager) IsUserAuthorized(jwtString string) error {
	if jwtString == "" {
		return errors.New("unauthorized")
	}
	token, err := jwt.ParseWithClaims(jwtString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secret), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("unauthorized")
	}

	_, ok := token.Claims.(*Claims)
	if !ok {
		return errors.New("invalid token claims")
	}
	return nil
}
