package listener

import (
	"fmt"
	"log"
	"net"
)

func GetNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return lis
}