package main

import (
	"log"
	"net"

	"google.golang.org/protobuf/proto"

	pb "Week-4/telemetry"
	network "network/packet"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}

			if conn == nil {
				return
			}

			serv := network.NewStream(1024)

			serv.OnError(func(err network.IOError) {
				conn.Close()
			})

			serv.SetConnection(conn)

			go func() {
				for msg := range serv.Incoming {
					log.Print(msg.Data)
					out := &pb.Telemetry{}
					err := proto.Unmarshal(msg.Data, out)
					if err != nil {
						log.Fatalln("Failed to decode", err)
					}
					log.Print(out)

					serv.Outgoing <- network.New(0, []byte("Recieved"))
				}
			}()
		}
	}()
}
