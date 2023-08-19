//go:generate go run go.uber.org/mock/mockgen -destination ../../testing/mocks/$GOFILE -package mocks . WalletRepository

package walletrepository

import (
	"context"
	"net/http"

	walletApi "buf.build/gen/go/ride/wallet/connectrpc/go/ride/wallet/v1alpha1/walletv1alpha1connect"
	pb "buf.build/gen/go/ride/wallet/protocolbuffers/go/ride/wallet/v1alpha1"
	"connectrpc.com/connect"
	"github.com/ride-app/marketplace-service/config"
	"github.com/sirupsen/logrus"
)

type WalletRepository interface {
	GetWallet(ctx context.Context, id string, authToken string) (*pb.Wallet, error)
}

type Impl struct {
	walletApi walletApi.WalletServiceClient
}

func New(config *config.Config) (*Impl, error) {
	logrus.Debug("Wallet Service Host: ", config.Wallet_Service_Host)
	client := walletApi.NewWalletServiceClient(
		http.DefaultClient,
		config.Wallet_Service_Host,
	)

	logrus.Info("Wallet Repository initialized")
	return &Impl{walletApi: client}, nil
}

func (r *Impl) GetWallet(ctx context.Context, id string, authToken string) (*pb.Wallet, error) {
	logrus.Info("Getting wallet from wallet service")
	req := connect.NewRequest(&pb.GetWalletRequest{
		Name: "users/" + id + "/wallet",
	})
	req.Header().Add("Authorization", authToken)

	res, err := r.walletApi.GetWallet(ctx, req)

	if connect.CodeOf(err) == connect.CodeNotFound {
		logrus.WithError(err).Error("Wallet not found")
		return nil, nil
	} else if err != nil {
		logrus.WithError(err).Error("Error getting wallet from wallet service")
		return nil, err
	}

	return res.Msg.Wallet, nil
}
