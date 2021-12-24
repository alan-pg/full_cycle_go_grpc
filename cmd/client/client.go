package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/alan-pg/fc2-grpc/pb/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to grpc server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	/* AddUser(client) */
	/* AddUserVerbose((client)) */
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Alan",
		Email: "alan@email.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make grpc request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Alan",
		Email: "alan@email.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not male GRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not make receive the msg: %v", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "a1",
			Name:  "Alan",
			Email: "alan@email.com",
		},
		&pb.User{
			Id:    "a2",
			Name:  "Alan2",
			Email: "alan2@email.com",
		},
		&pb.User{
			Id:    "a3",
			Name:  "Alan3",
			Email: "alan3@email.com",
		},
		&pb.User{
			Id:    "a4",
			Name:  "Alan4",
			Email: "alan4@email.com",
		},
		&pb.User{
			Id:    "a5",
			Name:  "Alan5",
			Email: "alan5@email.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating requests: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
