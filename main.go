package main

import (
	// "fmt"
	"net"
	"strings"
	"os"
	"log"

	pb "github.com/dansku/microservice_example_server/proto/calculate5"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {

	log.Println("=== Microservice Starting ===")

	// Load env file with server configuration

	err := godotenv.Load()
	if err != nil {
		grpclog.Fatal("Error loading .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")

	// ---------------------------------------

	listener, err := net.Listen("tcp", ":"+serverPort)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterCalculateServer(grpcServer, &server{})
	grpcServer.Serve(listener)

}

type server struct{}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {

	// Get values from request
	operation := strings.ToLower(request.Operation)

	// Test if operation is allowed
	allowedOperations := []string{"sum", "divide", "multiply", "subtract"}
	if !containInArray(allowedOperations,  operation) {
		response = &pb.Response{
			Number: 0,
		}
		return response, nil
	} 

	log.Printf("Performing %s \n", operation)
	log.Printf("Requested Values: %v \n", request.Numbers)

	var total float32 = 0
	total = CalculateValues(request.Numbers, operation)

	response = &pb.Response{
		Number: total,
	}

	log.Printf("Answer: %f \n", total)
	log.Printf("==============================")

	return response, nil
}

// CalculateValues get the slice of values and the type of operation and return the answer calculated
func CalculateValues(values []float32, operation string) float32 {

	var total float32
	operation = strings.ToLower(operation)

	switch operation {
	case "sum":
		total = 0
		for _, num := range values {
			total += num
		}
	case "subtract":
		total = 0
		for _, num := range values {
			total -= num
		}
	case "multiply":
		total = 1
		for _, num := range values {
			total = total * num
		}
	case "divide":
		total = 1
		for _, num := range values {
			total = total / num
		}
	default:
		total = 0
	}

	return total
}

// containInArray checks if the value exists in the slice
func containInArray(array []string, operation string) bool {
	for _, value := range array {
		if value == operation {
			return true
		}
	}

	return false
}