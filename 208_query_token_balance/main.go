package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
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

	abiData, err := os.ReadFile("./utils/erc20.abi")
	if err != nil {
		log.Fatalf("读取 ABI 文件失败: %v", err)
	}

	// 解析合约 ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiData)))
	if err != nil {
		log.Println("parse abi json error:", err)
	}

	tokenAddress := common.HexToAddress(utils.GetEnvParam("CONTRACT_ADDR"))

	accountAddress := common.HexToAddress(utils.GetEnvParam("ACCOUNT_2"))

	// 构造调用的 input data
	data, err := parsedABI.Pack("balanceOf", accountAddress)
	if err != nil {
		log.Println("pack abi error:", err)
	}

	// 构造调用消息
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}

	// 执行 eth_call
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Println("call contrace error:", err)
	}

	// 解析返回值
	var balance *big.Int
	err = parsedABI.UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		log.Println("parse return value error:", err)
	}

	// 打印余额（单位：最小单位，如 18 decimals）
	fmt.Printf("Token Balance: %s\n", balance.String())
}
