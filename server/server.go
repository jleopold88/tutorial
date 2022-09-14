package server

import (
	"database/sql"
	"fmt"
	"net"
	"tutorial/config"
	"tutorial/modules"

	proto "tutorial/prototype"

	_ "github.com/lib/pq"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewServer() {
	fmt.Println("Starting Server")

	cfg := config.NewConfig()

	listener, err := net.Listen("tcp", ":"+cfg.GRPCSERVER)
	if err != nil {
		errors.Wrap(err, "Server init failed")
		panic(err)
	}

	server := grpc.NewServer()

	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s"+
		" dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName)

	psql, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		errors.Wrap(err, "Database Connection failed")
		panic(err)
	}
	defer psql.Close()

	proto.RegisterTutorialServer(server, modules.NewRPC(modules.NewUseCase(modules.NewRepository(psql))))
	if err := server.Serve(listener); err != nil {
		errors.Wrap(err, "Failed to server the listener")
		panic(err)
	}

}
