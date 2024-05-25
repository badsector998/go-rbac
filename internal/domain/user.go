package domain

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint
	Name       string
	Email      string
	Password   string
	BurjoGroup BurjoGroup
}

type Claims struct {
	jwt.StandardClaims
	Role string
}
