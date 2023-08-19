package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ride-app/marketplace-service/api/gen/ride/marketplace/v1alpha1/marketplacev1alpha1connect"
	"github.com/ride-app/marketplace-service/api/interceptors"
	"github.com/ride-app/marketplace-service/config"
	"github.com/ride-app/marketplace-service/config/di"
	"github.com/ride-app/marketplace-service/utils/logger"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	config, err := config.New()

	log := logger.New(config)

	if err != nil {
		log.WithError(err).Fatal("Failed to read environment variables")
	}

	// Initialize service using dependency injection
	service, err := di.InitializeService(log, config)

	if err != nil {
		log.Fatalf("Failed to initialize service: %v", err)
	}

	log.Info("Service Initialized")

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

	// Create handler for RechargeService
	path, handler := marketplacev1alpha1connect.NewMarketplaceServiceHandler(service, connectInterceptors)

	// Create a new ServeMux and register the RechargeService handler
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	// Start the server and listen on the specified port
	// trunk-ignore(semgrep/go.lang.security.audit.net.use-tls.use-tls)
	panic(http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", config.Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	))
}
