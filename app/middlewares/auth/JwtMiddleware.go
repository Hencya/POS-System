package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"POS-System/helpers"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT   string
	ExpDuration int
}

func (cj *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(cj.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return c.JSON(http.StatusForbidden,
				helpers.BuildErrorResponse("failed to init token", http.StatusInternalServerError,
					e, nil))
		}),
	}
}

// GenerateToken jwt
func (cj *ConfigJWT) GenerateToken(Username string) string {
	claims := JwtCustomClaims{
		Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cj.ExpDuration))).Unix(),
		},
	}
	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(cj.SecretJWT))

	return token
}

//get user
func GetUser(c echo.Context) *JwtCustomClaims {
	medicalStaff := c.Get("user").(*jwt.Token)
	claims := medicalStaff.Claims.(*JwtCustomClaims)
	return claims
}
