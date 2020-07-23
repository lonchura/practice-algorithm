package token

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// implements PerRPCCredentials
type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"user":a.User, "password": a.Password}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var user string
	var password string

	if val, ok := md["user"]; ok { user = val[0] }
	if val, ok := md["password"]; ok { password = val[0] }

	// TODO rpc another auth service
	if !rpcCheckAuth(user, password) {
		// TODO grpc.Errorf deprecated
		return grpc.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}

// TODO rpc another auth service
func rpcCheckAuth(user string, password string) bool {
	if user == "lonchura" && password == "123456" {
		fmt.Println("auth success")
		return true
	}
	fmt.Println("auth failed")
	return false
}