package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	pb "github.com/madhusampangi/Say/saytext/proto"
)

func main() {
	port := flag.Int("p", 8080, "port to Listen")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Could not listen to port  %d : %v", *port, err)
	}
	log.Printf("Listening to port %d", *port)

	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("Could not serve : %v", err)
	}
}

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("Temp File creation failed %s: %v", f.Name(), err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("Could not close Temp File %s: %v", f.Name(), err)
	}

	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())

	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("Flite failed: %s", data)
	}
	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("Could not read Temp File %v", err)
	}
	return &pb.Speech{Audio: data}, nil
}
