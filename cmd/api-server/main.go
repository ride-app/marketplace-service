package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ride-app/marketplace-service/api/ride/marketplace/v1alpha1/v1alpha1connect"
	"github.com/ride-app/marketplace-service/config"
	"github.com/ride-app/marketplace-service/internal/api-handlers/interceptors"
	"github.com/ride-app/marketplace-service/internal/utils/logger"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	config, err := config.New()

	log := logger.New(config)

	if err != nil {
		log.WithError(err).Fatal("Failed to read environment variables")
	}

	// Create a context that, when cancelled, ends the JWKS background refresh goroutine.
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	panicInterceptor, err := interceptors.NewPanicInterceptor(ctx, log)

	if err != nil {
		log.Fatalf("Failed to initialize panic interceptor: %v", err)
	}

	authInterceptor, err := interceptors.NewAuthInterceptor(ctx, log)

	if err != nil {
		log.Fatalf("Failed to initialize auth interceptor: %v", err)
	}

	connectInterceptors := connect.WithInterceptors(panicInterceptor, authInterceptor)

	service, err := InitializeService(log, config)

	if err != nil {
		log.Fatalf("Failed to initialize service: %v", err)
	}

	log.Info("Service Initialized")

	path, handler := v1alpha1connect.NewMarketplaceServiceHandler(service, connectInterceptors)
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	// trunk-ignore(semgrep/go.lang.security.audit.net.use-tls.use-tls)
	panic(http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", config.Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	))

}
