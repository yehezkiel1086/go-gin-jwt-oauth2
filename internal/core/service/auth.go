package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/port"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/util"
	"golang.org/x/oauth2"
)

type AuthService struct {
	userRepo port.UserRepository
	oauthConf *oauth2.Config
	jwtConf *config.JWT
}

func InitAuthService(userRepo port.UserRepository, oauthConf *oauth2.Config, jwtConf *config.JWT) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		oauthConf: oauthConf,
		jwtConf: jwtConf,
	}
}

func (as *AuthService) Login(ctx context.Context, input *domain.User) (*domain.User, error) {
	// get user by email (check email)
	user, err := as.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return &domain.User{}, err
	}

	// compare input and user password
	if err := util.ComparePassword(user.Password, input.Password); err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (as *AuthService) GetGoogleLoginURL(ctx context.Context) string {
	return as.oauthConf.AuthCodeURL("state")
}

func (as *AuthService) GoogleCallback(ctx context.Context, code string) (string, int, error) {
	token, err := as.oauthConf.Exchange(ctx, code)
	if err != nil {
		return "", -1, err
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return "", -1, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var userInfo map[string]any
	json.Unmarshal(body, &userInfo)

	// Persist or update user
	user := &domain.User{
		Email:    userInfo["email"].(string),
		Name:     userInfo["name"].(string),
		Picture:  userInfo["picture"].(string),
		Provider: "google",
	}
	if _, err := as.userRepo.CreateOrUpdate(ctx, user); err != nil {
		return "", -1, fmt.Errorf("failed to create or update user: %w", err)
	}

	// Generate JWT
	jwtToken, duration, err := util.GenerateJWT(as.jwtConf, user)
	if err != nil {
		return "", -1, err
	}

	return jwtToken, duration, nil
}
