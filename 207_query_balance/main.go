package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hua-gao/ethclient/utils"
)

func main() {
	RPC_URL := utils.GetEnvParam("SEPOLIA_RPC_URL")
	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress(utils.GetEnvParam("ACCOUNT"))
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance: ", balance)

	blockId, err := strconv.ParseInt(utils.GetEnvParam("BLOCK_ID"), 10, 64)
	blockNumber := big.NewInt(blockId)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block balance:", balanceAt)
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("ethValue:", ethValue)
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("pendingBalance:", pendingBalance)
}
