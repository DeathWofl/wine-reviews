package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/deathwofl/wine-reviews/graph/model"
	storage "github.com/deathwofl/wine-reviews/pkg/storage/postgres"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var CurrentUserKey = "currentUser"

// AuthMiddleware function, authetication in middleware
func AuthMiddleware(serv storage.UserService) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				handler.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				handler.ServeHTTP(w, r)
				return
			}

			user, err := serv.User(claims["jwi"].(uint))
			if err != nil {
				handler.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CurrentUserKey, user)
			handler.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}

func bearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	authHeaderExtractor := &request.PostExtractionFilter{
		Filter:    bearerPrefixFromToken,
		Extractor: request.HeaderExtractor{"Authorization"},
	}

	authExtractor := &request.MultiExtractor{
		authHeaderExtractor,
		request.ArgumentExtractor{"access_token"},
	}
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv("JWT_SECRET"))
		return t, nil
	})

	return jwtToken, errors.Wrap(err, "parseToken error: ")
}

func GetCurrentUserFromContext(ctx context.Context) (*model.User, error) {
	if ctx.Value(CurrentUserKey) == nil {
		return nil, errors.New("No user in context")
	}

	user, ok := ctx.Value(CurrentUserKey).(*model.User)
	if !ok || user.ID == 0 {
		return nil, errors.New("no user")
	}

	return user, nil
}

// HashPassword generate hash password using hash256
func HashPassword(password string, u *model.User) error {
	bytePassword := []byte(password)

	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(passwordHash)

	return nil
}

func GenerateToken(u *model.User) (*model.AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 72)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "Wine-reviews",
	})

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &model.AuthToken{
		Token:     accessToken,
		ExpiredAt: expiredAt,
	}, nil
}

func ComparePassword(password string, u *model.User) error {
	passwordHashed := []byte(u.Password)
	log.Println(passwordHashed)
	log.Println(u.Password)
	log.Println([]byte(password))
	log.Println(password)
	return bcrypt.CompareHashAndPassword(passwordHashed, []byte(password))
}
