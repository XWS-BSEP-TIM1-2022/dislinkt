package services

import (
	connectionService "github.com/XWS-BSEP-TIM1-2022/dislinkt/util/proto/connection"
	jobService "github.com/XWS-BSEP-TIM1-2022/dislinkt/util/proto/job"
	messageService "github.com/XWS-BSEP-TIM1-2022/dislinkt/util/proto/message"
	postService "github.com/XWS-BSEP-TIM1-2022/dislinkt/util/proto/post"
	userService "github.com/XWS-BSEP-TIM1-2022/dislinkt/util/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewUserClient(address string) userService.UserServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return userService.NewUserServiceClient(conn)
}

func NewPostClient(address string) postService.PostServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Post service: %v", err)
	}
	return postService.NewPostServiceClient(conn)
}

func NewConnectionClient(address string) connectionService.ConnectionServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Connection service: %v", err)
	}
	return connectionService.NewConnectionServiceClient(conn)
}

func NewJobClient(address string) jobService.JobServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Connection service: %v", err)
	}
	return jobService.NewJobServiceClient(conn)
}

func NewMessageClient(address string) messageService.MessageServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return messageService.NewMessageServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
