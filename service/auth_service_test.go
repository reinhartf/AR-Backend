package service

import (
	gcontext "github.com/reinhartf/AR-Backend/context"
	"github.com/reinhartf/AR-Backend/model"
	"testing"
)

var (
	authService *AuthService
)

func init() {
	config := gcontext.LoadConfig("../")
	log := NewLogger(config)
	authService = NewAuthService(config, log)
}

func TestSignJWT(t *testing.T) {
	user := &model.User{
		ID:       "1",
		Email:    "test@1.com",
		Password: "123456",
	}
	tokenString, err := authService.SignJWT(user)
	if err != nil {
		t.Errorf("Error during signing JWT")
	}
	if *tokenString == "" || tokenString == nil {
		t.Errorf("Invalid JWT")
	}

}
