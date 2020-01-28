package rpc

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/params"
	"github.com/tidwall/pretty"
)



func TestRPCService_Discovery(t *testing.T) {

	server := NewServer()
	server.idgen = sequentialIDGenerator()
	if err := server.RegisterName("chainconfig", params.ClassicChainConfig); err != nil {
	//if err := server.RegisterName("chainconfig", (ctypes.ChainConfigurator)(generic.GenericCC{})); err != nil {
		t.Fatal(err)
	}

	defer server.Stop()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal("can't listen:", err)
	}

	defer listener.Close()
	go server.ServeListener(listener)

	var (
		deadlineDelta = 10 * time.Second
	)

	requests := []struct{
		method string
		params string
	}{
		{
		"rpc_modules", "",
		},
		{
			"rpc_discovery", "",
		},
	}

	// ...
	for i, request := range requests {
		conn, err := net.Dial("tcp", listener.Addr().String())
		if err != nil {
			t.Fatal("dial", err)
		}
		defer conn.Close()
		conn.SetDeadline(time.Now().Add(deadlineDelta))

		// Write req.
		conn.Write([]byte(fmt.Sprintf(`{"jsonrpc":"2.0","id":%d,"method":"%s", "params": [%s]}
`, i+1, request.method, request.params)))
		conn.(*net.TCPConn).CloseWrite()

		buff := make([]byte, 1024*1024)
		n, err := conn.Read(buff)
		if err != nil {
			t.Fatal("read conn", err)
		}
		//ret := gjson.ParseBytes(buff[:n])

		p := pretty.Pretty(buff[:n])
		log.Println(string(p))
	}
	//t.Fatal("jkjk")
}