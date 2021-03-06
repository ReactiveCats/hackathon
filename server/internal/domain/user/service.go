package user

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"server/internal/config"
	"server/internal/domain"
	"server/internal/ent"
	userent "server/internal/ent/user"
	"server/internal/platform"
)

type Service struct {
	client    *ent.UserClient
	jwtConfig config.Jwt
}

func NewService(client *ent.Client, config config.Config) *Service {
	return &Service{
		client:    client.User,
		jwtConfig: config.Jwt,
	}
}

func (s Service) JWTToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": user.ID})
	return token.SignedString([]byte(s.jwtConfig.Secret))
}

func (s Service) DataFromJWT(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong signing method")
		}
		return []byte(s.jwtConfig.Secret), nil
	})
	if err != nil {
		return 0, platform.WrapInternal(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, platform.WrapInternal(err)
	}
	if _, ok := claims["id"]; !ok {
		return 0, platform.NotFound("token payload entry not found")
	}

	return int(claims["id"].(float64)), nil
}

func (s Service) Login(ctx context.Context, username string) (string, error) {
	user, err := s.client.Query().
		Where(userent.Username(username)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", platform.NotFound("username not found")
		}
		return "", platform.WrapInternal(err)
	}

	return s.JWTToken(domain.UserFromEnt(user))
}

func (s Service) Signup(ctx context.Context, username string) (string, error) {
	user, err := s.client.Create().
		SetUsername(username).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return "", platform.Conflict("username already exists")
		}
		return "", platform.WrapInternal(err)
	}

	return s.JWTToken(domain.UserFromEnt(user))
}

func (s Service) ByID(ctx context.Context, dto domain.GetUserDTO) (*domain.User, error) {
	usr, err := s.client.Query().
		Where(userent.ID(dto.ID)).
		Only(ctx)
	if err != nil {
		return nil, platform.WrapInternal(err)
	}

	return domain.UserFromEnt(usr), nil
}

func (s Service) ByUsername(ctx context.Context, username string) (*domain.User, error) {
	usr, err := s.client.Query().
		Where(userent.Username(username)).
		WithTasks().
		Only(ctx)
	if err != nil {
		return nil, platform.WrapInternal(err)
	}

	return domain.UserFromEnt(usr), nil
}
