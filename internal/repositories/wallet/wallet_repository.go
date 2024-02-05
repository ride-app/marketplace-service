//go:generate go run go.uber.org/mock/mockgen -destination ./mock/$GOFILE . WalletRepository

package walletrepository

import (
	"context"
	"net/http"

	walletApi "buf.build/gen/go/ride/payments/connectrpc/go/ride/payments/v1alpha1/paymentsv1alpha1connect"
	pb "buf.build/gen/go/ride/payments/protocolbuffers/go/ride/payments/v1alpha1"
	"connectrpc.com/connect"
	"github.com/deb-tech-n-sol/go/pkg/logger"
	"github.com/ride-app/marketplace-service/config"
)

type WalletRepository interface {
	GetWallet(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Wallet, error)
}

type Impl struct {
	walletApi walletApi.PaymentsServiceClient
}

func New(log logger.Logger, config *config.Config) (*Impl, error) {
	log.Debug("wallet Service Host: ", config.Payment_Service_Host)
	client := walletApi.NewPaymentsServiceClient(
		http.DefaultClient,
		config.Payment_Service_Host,
	)

	log.Info("wallet Repository initialized")
	return &Impl{walletApi: client}, nil
}

func (r *Impl) GetWallet(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Wallet, error) {
	log.Info("getting wallet from wallet service")
	req := connect.NewRequest(&pb.GetWalletRequest{
		Name: "users/" + id + "/wallet",
	})
	req.Header().Add("Authorization", authToken)

	res, err := r.walletApi.GetWallet(ctx, req)

	if connect.CodeOf(err) == connect.CodeNotFound {
		log.WithError(err).Error("wallet not found")
		return nil, nil
	} else if err != nil {
		log.WithError(err).Error("error getting wallet from wallet service")
		return nil, err
	}

	return res.Msg.Wallet, nil
}
