package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ride-app/marketplace-service/api/gen/ride/marketplace/v1alpha1/marketplacev1alpha1connect"
	"github.com/ride-app/marketplace-service/config"
	"github.com/ride-app/marketplace-service/di"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/sirupsen/logrus"
)

func main() {
	service, err := di.InitializeService()

	if err != nil {
		logrus.Fatalf("Failed to initialize service: %v", err)
	}

	logrus.Info("Service Initialized")

	path, handler := marketplacev1alpha1connect.NewMarketplaceServiceHandler(service)
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	panic(http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", config.Env.Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	))

}

func init() {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	})

	logrus.SetLevel(logrus.InfoLevel)

	err := cleanenv.ReadEnv(&config.Env)

	if config.Env.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if err != nil {
		logrus.Warnf("Could not load config: %v", err)
	}
}
