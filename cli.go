package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// CLI processes command line related operations
type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("	getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("	createblockchain -address ADDRESS - Create a new blockchain and send the genesis block reward to ADDRESS")
	fmt.Println("	show - Show all the blocks of the blockchain")
	fmt.Println("	send -from FROM -to TO -amount AMOUNT - Send AMOUNT from FROM address to TO address")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses CLI arguments and runs commands
func (cli *CLI) Run() {
	cli.validateArgs()

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	showCmd := flag.NewFlagSet("show", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get the balance for")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send the genesis block reward to")
	sendFromAddress := sendCmd.String("from", "", "Sender's wallet address")
	sendToAddress := sendCmd.String("to", "", "Receiver's wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount of coins to send")

	switch os.Args[1] {
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "show":
		err := showCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	if showCmd.Parsed() {
		cli.showBlockchain()
	}

	if sendCmd.Parsed() {
		if *sendFromAddress == "" || *sendToAddress == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}
		cli.send(*sendFromAddress, *sendToAddress, *sendAmount)
	}
}

func (cli *CLI) showBlockchain() {
	bc := NewBlockchain("")
	defer bc.db.Close()
	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) createBlockchain(address string) {
	bc := CreateBlockchain(address)
	bc.db.Close()
	fmt.Println("Blockchain created!")
}

func (cli *CLI) getBalance(address string) {
	bc := NewBlockchain(address)
	defer bc.db.Close()

	balance := 0
	UTXOs := bc.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("%s's balance is %d\n", address, balance)
}

func (cli *CLI) send(from, to string, amount int) {
	bc := NewBlockchain(from)
	defer bc.db.Close()

	tx := NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*Transaction{tx})
	fmt.Println("Transaction sent!")
}
