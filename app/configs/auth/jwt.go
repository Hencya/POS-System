package auth

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"POS-System/app/middlewares/auth"
)

func SetupJwt() auth.ConfigJWT {
	_ = godotenv.Load()

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExp := os.Getenv("JWT_EXPIRE")

	Exp, _ := strconv.Atoi(jwtExp)
	configJWT := auth.ConfigJWT{
		SecretJWT:   jwtSecret,
		ExpDuration: Exp,
	}

	return configJWT
}
