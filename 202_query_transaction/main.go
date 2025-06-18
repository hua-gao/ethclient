package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hua-gao/ethclient/utils"
)

func main() {
	RPC_URL := utils.GetEnvParam("SEPOLIA_RPC_URL")

	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	// 获取链的 ChainID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 获取区块信息
	blockId, err := strconv.ParseInt(utils.GetEnvParam("BLOCK_ID"), 10, 64)
	blockNumber := big.NewInt(blockId)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())
		fmt.Println(tx.Data())
		fmt.Println(tx.To().Hex())

		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex())
		} else {
			log.Println("get sender error:", err)
		}
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Println("get receipt error:", err)
		}
		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		break
	}

	blockHash := common.HexToHash(utils.GetEnvParam("TRANS_ID"))
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Println("get transction count error:", err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Println("get transation in block error:", err)
		}

		fmt.Println(tx.Hash().Hex())
		break
	}

	txHash := common.HexToHash(utils.GetEnvParam("TRANS_ID"))
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Println("get transation by hash error:", err)
	}
	fmt.Println(isPending)
	fmt.Println(tx.Hash().Hex())
}
