//go:generate go run go.uber.org/mock/mockgen -destination ../../../pkg/testing/mocks/$GOFILE -package mocks . WalletRepository

package walletrepository

import (
	"context"
	"net/http"

	walletApi "buf.build/gen/go/ride/wallet/connectrpc/go/ride/wallet/v1alpha1/walletv1alpha1connect"
	pb "buf.build/gen/go/ride/wallet/protocolbuffers/go/ride/wallet/v1alpha1"
	"connectrpc.com/connect"
	"github.com/ride-app/marketplace-service/config"
	"github.com/ride-app/marketplace-service/internal/utils/logger"
)

type WalletRepository interface {
	GetWallet(ctx context.Context, log logger.Logger, id string, authToken string) (*pb.Wallet, error)
}

type Impl struct {
	walletApi walletApi.WalletServiceClient
}

func New(log logger.Logger, config *config.Config) (*Impl, error) {
	log.Debug("wallet Service Host: ", config.Wallet_Service_Host)
	client := walletApi.NewWalletServiceClient(
		http.DefaultClient,
		config.Wallet_Service_Host,
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
