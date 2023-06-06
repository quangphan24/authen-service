package grpc

import (
	"authen-service/proto/authen"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const JWT_SECRET_KEY = "GO_BASE"

type ServerGRPC struct {
	authen.UnimplementedAuthenServiceServer
}

func (s *ServerGRPC) VerifyToken(c context.Context, in *authen.String) (*authen.Bool, error) {
	if in.GetValue() == "" {
		return nil, errors.New("You need access permission")
	}

	newToken, err := jwt.Parse(in.GetValue(), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method:%v", token.Header["alg"]))
		}
		signature := []byte(JWT_SECRET_KEY)
		return signature, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := newToken.Claims.(jwt.MapClaims)
	if !ok || !newToken.Valid {
		return nil, errors.New("Couldn't parse claims")
	}

	//check expired time
	if int64(claims["exp_at"].(float64)) < time.Now().Local().Unix() {
		return nil, errors.New("Token expired")
	}
	return &authen.Bool{Value: true}, nil
}
