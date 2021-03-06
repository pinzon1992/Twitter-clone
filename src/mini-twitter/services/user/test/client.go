package test

import (
	"google.golang.org/grpc"
	"log"
	"mini-twitter/services/user/userpb"
	"os"
)

type UserClient struct {
	UserDB   userpb.UserServiceClient
	User     *grpc.ClientConn
	UserPort string
}

var uc UserClient

func InitializeUserClient() {
	var err error
	uc.UserPort = os.Getenv("USER_PORT")
	log.Println("userPort =", uc.UserPort)
	uc.User, err = grpc.Dial("localhost:"+uc.UserPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	uc.UserDB = userpb.NewUserServiceClient(uc.User)

	log.Println("User Client created at port =", uc.UserPort)
}
