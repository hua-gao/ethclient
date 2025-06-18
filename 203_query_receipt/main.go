package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hua-gao/ethclient/utils"
)

func main() {
	RPC_URL := utils.GetEnvParam("SEPOLIA_RPC_URL")

	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	blockId, err := strconv.ParseInt(utils.GetEnvParam("BLOCK_ID"), 10, 64)
	blockNumber := big.NewInt(blockId)
	blockHash := common.HexToHash(utils.GetEnvParam("BLOCK_HASH"))

	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Println("get receipt err: ", err)
	}

	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Println("get receipt err: ", err)
	}
	fmt.Println(receiptByHash[0] == receiptsByNum[0]) // true

	for _, receipt := range receiptByHash {
		fmt.Println("Status:", receipt.Status)
		fmt.Println("Logs:", receipt.Logs)
		fmt.Println("TxHash:", receipt.TxHash.Hex())
		fmt.Println("TransactionIndex:", receipt.TransactionIndex)
		fmt.Println("ContractAddress:", receipt.ContractAddress.Hex())
		break
	}

	txHash := common.HexToHash(utils.GetEnvParam("BLOCK_HASH_2"))
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Println("get transactionReceipt error:", err)
	}
	fmt.Println(receipt.Status)
	fmt.Println(receipt.Logs)
	fmt.Println(receipt.TxHash.Hex())
	fmt.Println(receipt.TransactionIndex)
	fmt.Println(receipt.ContractAddress.Hex())
}
