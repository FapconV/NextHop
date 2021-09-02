package main

import (
	"log"
	"net"

	"google.golang.org/protobuf/proto"

	pb "Week-4/telemetry"

	network "network/packet"
)

const (
	address = "localhost:50051"
)

func main() {
	var cid uint64 = 1
	bp := "tcp"
	sid := "venu"
	mv := "1.2"
	var csT uint64 = 7
	var mts uint64 = 5
	var ts uint64 = 2
	n := "venu"
	agd := true
	//vbt := pb.TelemetryField_BoolValue{
	//			BoolValue : true,
	//		}
	var ced uint64 = 5

	c := pb.Telemetry{
		CollectionId:           &cid,
		BasePath:               &bp,
		SubscriptionIdentifier: &sid,
		ModelVersion:           &mv,
		CollectionStartTime:    &csT,
		MsgTimestamp:           &mts,
		Fields: []*pb.TelemetryField{
			{
				Timestamp:   &ts,
				Name:        &n,
				AugmentData: &agd,
				//ValueByType: &vbt,
				Fields: []*pb.TelemetryField{},
			},
		},
		CollectionEndTime: &ced,
	}

	// Contact the server and print out its response.
	msg, err := proto.Marshal(&c)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	log.Print(msg)

	// Set up a connection to the server.
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	s := network.NewStream(1024)
	s.OnError(func(err network.IOError) {
		conn.Close()
	})

	s.SetConnection(conn)

	s.Outgoing <- network.New(0, []byte(msg))

	recv := <-s.Incoming

	log.Print(recv.Data)

}
