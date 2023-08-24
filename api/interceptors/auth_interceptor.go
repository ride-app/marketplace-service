package interceptors

import (
	"context"
	"errors"
	"time"

	"connectrpc.com/connect"
	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ride-app/marketplace-service/utils/logger"
)

func NewAuthInterceptor(ctx context.Context, log logger.Logger) (*connect.UnaryInterceptorFunc, error) {
	jwksURI := "https://www.googleapis.com/service_accounts/v1/jwk/securetoken@system.gserviceaccount.com"

	options := keyfunc.Options{
		Ctx: ctx,
		RefreshErrorHandler: func(err error) {
			formattedErrMsg := logger.FormatString("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
			log.Fatalf(formattedErrMsg)
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	jwks, err := keyfunc.Get(jwksURI, options)

	if err != nil {
		return nil, err
	}

	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
   			if req.Header().Get("authorization") == "" {
   				formattedErrMsg := logger.FormatString("no token provided")
   				return nil, connect.NewError(
   					connect.CodeUnauthenticated,
   					errors.New(formattedErrMsg),
   				)
   			}

   			if req.Header().Get("authorization")[:7] != "Bearer " {
   				formattedErrMsg := logger.FormatString("invalid token format")
   				return nil, connect.NewError(
   					connect.CodeUnauthenticated,
   					errors.New(formattedErrMsg),
   				)
   			}
			token, err := jwt.Parse(req.Header().Get("authorization")[7:], jwks.Keyfunc)

   			if !token.Valid {
   				formattedErrMsg := logger.FormatString("invalid token")
   				return nil, connect.NewError(
   					connect.CodeUnauthenticated,
   					errors.New(formattedErrMsg),
   				)
   			}

			req.Header().Set("uid", token.Claims.(jwt.MapClaims)["user_id"].(string))

   			if err != nil {
   				formattedErrMsg := logger.FormatString(err.Error())
   				return nil, connect.NewError(
   					connect.CodeUnauthenticated,
   					errors.New(formattedErrMsg),
   				)
   			}

			return next(ctx, req)
		})
	}

	interceptorFunc := connect.UnaryInterceptorFunc(interceptor)

	return &interceptorFunc, nil
}
