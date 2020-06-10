package service

import (
	"context"
	"database/sql"
	"strconv"

	"AliveVirtualGift_AccountService/src/proto"
	"AliveVirtualGift_AccountService/src/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var clientSessionService proto.SessionServiceClient

//DialToServiceServer ...
func DialToServiceServer(port int) {

	conn, err := grpc.Dial("localhost:"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	clientSessionService = proto.NewSessionServiceClient(conn)
}

//serviceServer ...
type serviceServer struct {
	db *sql.DB
}

//NewAccountServiceServer ...
func NewAccountServiceServer(db *sql.DB) proto.AccountServiceServer {
	return &serviceServer{db: db}
}

func (s *serviceServer) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {

	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT Password, Id, Type FROM Account WHERE Username = ?", request.GetUsername())
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to select Username from Account-> "+err.Error())
	}
	defer rows.Close()

	var hashedPassoword string
	var id uint64
	var accType []uint8

	for rows.Next() {
		err = rows.Scan(&hashedPassoword, &id, &loginType)
		if err != nil {
			return nil, err
		}
	}

	err := utils.ComparePasswords(hashedPassoword, request.GetPassword())
	if err != nil {
		return nil, err
	}

	tokenString, err := clientSessionService.CreateToken(ctx, &proto.Payload{
		Id: id,
		Type: proto.Type(proto.Type_value[string(accType)])
	})

	return &proto.LoginResponse{
		Token: &tokenString.Token,
	}, nil
}

func (s *serviceServer) Logout(ctx context.Context, request *proto.LogoutRequest) (*proto.Status, error) {

	status, err := clientSessionService.DeleteToken(ctx, &proto.TokenString{
		Token: &request.GetToken(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.Status{
		Success: true,
	}, nil
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
