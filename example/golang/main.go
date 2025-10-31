package main

import (
	"context"
	"log"
	"math/big"

	"github.com/CycleNetwork-Labs/cycle-unit-withdraw-sdk/withdraw-go/transactions"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()

	sender := common.HexToAddress("")
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := client.PendingNonceAt(ctx, sender)
	gasPrice, err := client.SuggestGasPrice(ctx)
	gasLimit := uint64(220000)
	sourceChainID := int64(11155111)
	destinationChainID := int64(421614)
	destinationAddress := common.HexToAddress("0x0")

	params := transactions.WithdrawParams{
		Client:             client,
		Nonce:              nonce,
		GasPrice:           gasPrice,
		GasLimit:           gasLimit,
		TokenSymbol:        "USDT",
		SourceChainID:      sourceChainID,
		DestinationChainID: destinationChainID,
		DestinationAddress: destinationAddress,
		Amount:             big.NewInt(1000000),
	}
	tx, err := transactions.Build(params)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(tx)
}
