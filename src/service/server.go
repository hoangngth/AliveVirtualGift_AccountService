package service

import (
	"context"
	"database/sql"
	"fmt"

	"AliveVirtualGift_AccountService/src/proto"
	"AliveVirtualGift_AccountService/src/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var clientSessionService proto.SessionServiceClient

//DialToServiceServer ...
func DialToServiceServer(port string) {

	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
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
		err = rows.Scan(&hashedPassoword, &id, &accType)
		if err != nil {
			return nil, err
		}
	}

	pwdErr := utils.ComparePasswords(hashedPassoword, request.GetPassword())
	if pwdErr != nil {
		return nil, pwdErr
	}

	tokenString, err := clientSessionService.CreateToken(ctx, &proto.AccountInfo{
		Id:   id,
		Type: proto.Type(proto.Type_value[string(accType)]),
	})

	return &proto.LoginResponse{
		Token: tokenString.Token,
	}, nil
}

func (s *serviceServer) Logout(ctx context.Context, request *proto.LogoutRequest) (*proto.LogoutResponse, error) {

	stat, err := clientSessionService.DeleteToken(ctx, &proto.TokenString{
		Token: request.GetToken(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.LogoutResponse{
		IsLoggedOut: stat.Success,
	}, nil
}

func (s *serviceServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {

	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	hashedPassword := utils.HashAndSalt(request.GetPassword())
	res, err := conn.ExecContext(ctx, "INSERT INTO Account (Name, Email, Username, Password) VALUES (?, ?, ?, ?)",
		request.GetName(), request.GetEmail(), request.GetUsername(), hashedPassword)

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to insert into Account-> "+err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to retrieve id for created Account-> "+err.Error())
	}
	_ = id

	return &proto.CreateResponse{
		IsCreated: true,
	}, nil
}

func (s *serviceServer) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {

	stat, err := clientSessionService.CheckToken(ctx, &proto.TokenString{
		Token: request.GetToken(),
	})
	if err != nil {
		return nil, err
	}

	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	accountID, err := clientSessionService.GetAccountIDFromToken(ctx, &proto.TokenString{
		Token: request.GetToken(),
	})
	id := &accountID.Id
	hasedPassword := utils.HashAndSalt(request.GetPassword())

	res, err := conn.ExecContext(ctx, "UPDATE Account SET Name = ?, Email = ?, Avatar = ?, Password = ? WHERE  Id = ?",
		request.GetName(), request.GetEmail(), request.GetAvatar(), hasedPassword, id)

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to update Account-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Account with ID='%d' is not found or unchanged", accountID))
	}

	return &proto.UpdateResponse{
		IsUpdated: stat.Success,
	}, nil
}

func (s *serviceServer) Show(ctx context.Context, request *proto.ShowRequest) (*proto.ShowResponse, error) {

	_, err := clientSessionService.CheckToken(ctx, &proto.TokenString{
		Token: request.GetToken(),
	})
	if err != nil {
		return nil, err
	}

	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	accountID, err := clientSessionService.GetAccountIDFromToken(ctx, &proto.TokenString{
		Token: request.GetToken(),
	})
	id := &accountID.Id

	rows, err := conn.QueryContext(ctx, "SELECT Name, Email, Avatar, Diamond FROM Account WHERE Id = ?", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to select from Account-> "+err.Error())
	}
	defer rows.Close()

	var showResponse proto.ShowResponse
	for rows.Next() {
		err = rows.Scan(&showResponse.Name, &showResponse.Email, &showResponse.Avatar, &showResponse.Diamond)
		if err != nil {
			return nil, err
		}
	}

	return &showResponse, nil
}
