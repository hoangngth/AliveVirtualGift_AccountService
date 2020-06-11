package main

import (
	"AliveVirtualGift_AccountService/src/database"
	"AliveVirtualGift_AccountService/src/proto"
	"AliveVirtualGift_AccountService/src/service"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {

	// Load environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	listener, err := net.Listen("tcp", ":"+os.Getenv("ACCOUNT_SERVICE_PORT"))
	if err != nil {
		panic(err)
	}
	log.Print("Account Service running on port :" + os.Getenv("ACCOUNT_SERVICE_PORT"))

	var newdb database.DBInfo
	db, err := newdb.GetDB()
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
		return
	}

	// Dial to Session Service
	service.DialToServiceServer(os.Getenv("SESSION_SERVICE_PORT"))

	srv := grpc.NewServer()
	service := service.NewAccountServiceServer(db)

	proto.RegisterAccountServiceServer(srv, service)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx := context.Background()
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("Shutting down gRPC Account service server...")

			srv.GracefulStop()

			<-ctx.Done()
		}
	}()

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}
