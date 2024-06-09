package interceptors

import (
	"context"
	"encoding/json"
	"errors"

	"connectrpc.com/connect"
	services "github.com/imkrish7/ecoship-api/services/jwt"
)

func AuthInterceptors() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Header().Get("tokenHeader") == "" {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("Unauthorized"),
				)
			} else {
				token := req.Header().Get("tokenHeader")[7:]
				authToken, err := services.VerifyToken(token)
				if err != nil {
					return nil, connect.NewError(
						connect.CodeUnauthenticated,
						err,
					)
				}
				if authToken == nil {
					return nil, connect.NewError(
						connect.CodeUnauthenticated,
						errors.New("Unauthorized"),
					)
				}

				var newClaims services.UserClaims
				json.Unmarshal(authToken.Claims(), &newClaims)
				ctx = context.WithValue(ctx, "userdata", newClaims)
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
