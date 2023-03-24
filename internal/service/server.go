package service

import (
	"git.miem.hse.ru/1206/app"
	"git.miem.hse.ru/1206/app/errs"
	"git.miem.hse.ru/1206/app/middleware"
	"git.miem.hse.ru/1206/material-library/internal/domain"
	"git.miem.hse.ru/1206/proto/pb"
	"google.golang.org/grpc"
)

type LibraryServer struct {
	pb.UnimplementedLibraryServer
	uc domain.UseCase
}

func NewLibraryServer(cfg *app.GRPCConfig, uc domain.UseCase) (*app.GRPCServer, error) {
	var opts []grpc.ServerOption
	opts = append(opts,
		grpc.ChainUnaryInterceptor(
			middleware.InterceptorRecovery(),
			middleware.InterceptorLogger(),
		),
	)

	grpcServer, err := app.NewGRPCServer(cfg, opts...)
	pb.RegisterLibraryServer(grpcServer.Ser, &LibraryServer{uc: uc})
	if err != nil {
		return nil, errs.New(err)
	}
	grpcServer.Run()
	return grpcServer, nil
}
