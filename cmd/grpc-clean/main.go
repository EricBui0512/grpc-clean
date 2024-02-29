package main

import (
	"fmt"
	"log"
	"net"

	interfaces "github.com/EricBui0512/grpc-clean/pkg/v1"
	repo "github.com/EricBui0512/grpc-clean/pkg/v1/repository"
	usecase "github.com/EricBui0512/grpc-clean/pkg/v1/usecase"
	"gorm.io/gorm"

	dbConfig "github.com/EricBui0512/grpc-clean/internal/db"
	"github.com/EricBui0512/grpc-clean/internal/models"
	handler "github.com/EricBui0512/grpc-clean/pkg/v1/handler/grpc"
	"google.golang.org/grpc"
)

func main() {

	db := dbConfig.DbConn()
	migrations(db)

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("ERROT STARTING THE SERVER : %v", err)
	}

	grpcServer := grpc.NewServer()

	userUseCase := initUserServer(db)
	handler.NewServer(grpcServer, userUseCase)

	log.Fatal(grpcServer.Serve(lis))

}

func initUserServer(db *gorm.DB) interfaces.UseCaseInterface {
	userRepo := repo.New(db)
	return usecase.New(userRepo)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
