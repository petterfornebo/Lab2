package client

import (
	"context"
	"flag"
	"log"

	"github.com/dat520-2021/petterfornebo-labs/tree/master/lab2/grpc/proto"
	"google.golang.org/grpc"
)

var serverAddr = flag.String("s", "", "Server address")

func init() {
	flag.Parse()
}

func main() {
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure()) //dial the grpc connection, disable transport security for this Client connection
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close() // close when finish

	client := proto.NewKeyValueServiceClient(conn) //make a stud on the conn return from dial
	testValues := map[string]string{               //make som key/value pairs to test with
		"value 1": "",
		"value 2": "rule Britannia",
		"value 3": "Britannia rule the waves",
	}

	for k, v := range testValues {
		resp, err := client.Insert(context.Background(), &proto.InsertRequest{Key: k, Value: v})
		if err != nil {
			log.Fatalln(err)
		}
		if !resp.Success { // if success is not returned, log the error
			log.Fatalf("Inserting value for key %s failed (server reported success == false)", k)
		}
	}

	for k, v := range testValues {
		resp, err := client.Lookup(context.Background(), &proto.LookupRequest{Key: k})
		if err != nil {
			log.Fatalf("error looking up value for key %s: %v", k, err)
		}
		if resp.Value != v {
			log.Fatalf("wrong value for key %s: %v returned (expected %s)", k, resp.Value, v) // checking if the values at the server are the same as mine: MADS FJERNE?
		}
	}

	resp, err := client.Keys(context.Background(), &proto.KeysRequest{})
	if err != nil {
		log.Fatalln("error getting keys:", err)
	}
	if e, a := len(testValues), len(resp.Keys); e != a { //check if there are more keys at the server than in testvalues => something wrong
		log.Fatalf("Wrong number of keys returned. Expected %d, got %d", e, a)
	}
}
