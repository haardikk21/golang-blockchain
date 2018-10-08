package main

import (
	"fmt"
	"log"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)
	if len(minerAddress) > 0 {
		if ValidateAddress(minerAddress) {
			fmt.Printf("Mining... Rewards will be sent to %s\n", minerAddress)
		} else {
			log.Panic("Invalid miner address")
		}
	}

	StartServer(nodeID, minerAddress)
}
