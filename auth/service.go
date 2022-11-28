package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)

	//pakai jwt token karna nanti menggunakan method balikan dari jwt
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

//seharusnya tidak ditaruh disini
var SECRET_KEY = []byte("aldianstartup_s3cr3t_k3y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	//membuat payload jwt
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	//Algoritma signing method HS256 jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//membuat ttd jwt / VERIFY SIGNATURE
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	//parse parameter token dan func
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		//secret HS256 salah satu dari HMAC
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
