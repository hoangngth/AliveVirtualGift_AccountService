package service

import (
	"context"
	"database/sql"

	"AliveVirtualGift_AccountService/src/proto"
)

//serviceServer ...
type serviceServer struct {
	db *sql.DB
}

//NewAccountServiceServer ...
func NewAccountServiceServer(db *sql.DB) proto.AccountServiceServer {
	return &serviceServer{db: db}
}

func (s *serviceServer) Login(ctx context.Context, request *proto.LoginRequest) (*proto.Status, error) {
	return nil, nil
}

func (s *serviceServer) Logout(ctx context.Context, request *proto.LogoutRequest) (*proto.Status, error) {
	return nil, nil
}

func (s *serviceServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.Status, error) {
	return nil, nil
}

func (s *serviceServer) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.Status, error) {
	return nil, nil
}

func (s *serviceServer) Show(ctx context.Context, request *proto.ShowRequest) (*proto.Status, error) {
	return nil, nil
}
