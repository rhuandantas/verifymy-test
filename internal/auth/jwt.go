package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/rhuandantas/verifymy-test/internal/config"
	"net/http"
	"strings"
	"time"
)

//go:generate mockgen -source=$GOFILE -package=mock_auth -destination=../../test/mock/auth/$GOFILE

type Token interface {
	GenerateToken(email string) (string, error)
}

type JwtToken struct {
	config config.ConfigProvider
}

func NewJwtToken(config config.ConfigProvider) *JwtToken {
	return &JwtToken{
		config: config,
	}
}

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (jt *JwtToken) GenerateToken(email string) (string, error) {
	secret := []byte(jt.config.GetEnv("AUTH_SECRET"))
	claims := &jwtCustomClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return t, nil
}

func (jt *JwtToken) VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := jt.getToken(c)
		secret := []byte(jt.config.GetEnv("AUTH_SECRET"))
		tkn, err := jwt.ParseWithClaims(tokenStr, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		if !tkn.Valid {
			return c.JSON(http.StatusUnauthorized, "invalid")
		}

		return next(c)
	}
}

func (jt *JwtToken) getToken(c echo.Context) string {
	bearer := c.Request().Header.Get("Authorization")
	tokenStr := strings.Split(bearer, " ")[1]
	return tokenStr
}
