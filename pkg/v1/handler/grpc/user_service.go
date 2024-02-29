package grpc

import (
	"context"
	"errors"

	"github.com/EricBui0512/grpc-clean/internal/models"
	interfaces "github.com/EricBui0512/grpc-clean/pkg/v1"
	pb "github.com/EricBui0512/grpc-clean/proto"
	"google.golang.org/grpc"
)

type UserServStruct struct {
	useCase interfaces.UseCaseInterface
	pb.UnimplementedUserServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &UserServStruct{useCase: usecase}
	pb.RegisterUserServiceServer(grpcServer, userGrpc)
}

func (srv *UserServStruct) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {
	data := srv.transformUserRPC(req)

	if data.Email == "" || data.Name == "" {
		return &pb.UserProfileResponse{}, errors.New("please provide all fields")
	}

	user, err := srv.useCase.Create(data)

	if err != nil {
		return &pb.UserProfileResponse{}, err
	}

	return srv.transformUserModel(user), nil

}

func (srv *UserServStruct) transformUserRPC(req *pb.CreateUserRequest) models.User {
	return models.User{Name: req.GetName(), Email: req.GetEmail()}
}

func (srv *UserServStruct) transferUserModel(user models.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{Name: user.Name, Email: user.Email}
}
