package main

import "fmt"

func (cli *CLI) createBlockchain(address, nodeID string) {
	bc := CreateBlockchain(address, nodeID)
	defer bc.db.Close()

	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Blockchain created!")
}
