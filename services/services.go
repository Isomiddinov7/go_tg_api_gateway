package services

import (
	"go_tg_api_gateway/config"
	"go_tg_api_gateway/genproto/coins_service"
	"go_tg_api_gateway/genproto/users_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	CoinService() coins_service.CoinsServiceClient
	BuyOrSale() coins_service.BuyOrSellClient
	FileImage() coins_service.ImagesServiceClient
	UserService() users_service.UserServiceClient
	UserTransaction() users_service.UserSellOrBuyServiceClient
	Messages() users_service.UserMessageListClient
}

type grpcClients struct {
	coinService            coins_service.CoinsServiceClient
	userService            users_service.UserServiceClient
	buyorsaleService       coins_service.BuyOrSellClient
	fileimageService       coins_service.ImagesServiceClient
	usertransactionService users_service.UserSellOrBuyServiceClient
	messageService         users_service.UserMessageListClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	// Coin Service...
	connCoinService, err := grpc.NewClient(
		cfg.CoinServiceHost+cfg.CoinGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	//User Service...
	connUserService, err := grpc.NewClient(
		cfg.UserServiceHost+cfg.UserGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		coinService:            coins_service.NewCoinsServiceClient(connCoinService),
		buyorsaleService:       coins_service.NewBuyOrSellClient(connCoinService),
		fileimageService:       coins_service.NewImagesServiceClient(connCoinService),
		userService:            users_service.NewUserServiceClient(connUserService),
		usertransactionService: users_service.NewUserSellOrBuyServiceClient(connUserService),
		messageService:         users_service.NewUserMessageListClient(connUserService),
	}, nil
}

func (g *grpcClients) CoinService() coins_service.CoinsServiceClient {
	return g.coinService
}

func (g *grpcClients) BuyOrSale() coins_service.BuyOrSellClient {
	return g.buyorsaleService
}

func (g *grpcClients) FileImage() coins_service.ImagesServiceClient {
	return g.fileimageService
}

func (g *grpcClients) UserService() users_service.UserServiceClient {
	return g.userService
}

func (g *grpcClients) UserTransaction() users_service.UserSellOrBuyServiceClient {
	return g.usertransactionService
}

func (g *grpcClients) Messages() users_service.UserMessageListClient {
	return g.messageService
}
