package transactions

import (
	"math/big"

	"github.com/CycleNetwork-Labs/cycle-unit-withdraw-sdk/withdraw-go/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	TokenTypeNative = 1
	TokenTypeERC20  = 0
)

type WithdrawParams struct {
	Client             *ethclient.Client
	Nonce              uint64
	GasPrice           *big.Int
	GasLimit           uint64
	TokenSymbol        string
	SourceChainID      int64
	DestinationChainID int64
	DestinationAddress common.Address
	Amount             *big.Int
}

func Build(params WithdrawParams) (*types.Transaction, error) {
	tokenID, sourceToken, destinationToken, err := queryTokenInfo(
		params.TokenSymbol,
		params.SourceChainID,
		params.DestinationChainID,
	)
	if err != nil {
		return nil, err
	}

	if params.SourceChainID == params.DestinationChainID {
		switch sourceToken.NativeToken {
		case TokenTypeNative:
			return TransferNativeToken(
				params.Nonce,
				params.GasPrice,
				params.GasLimit,
				params.DestinationAddress,
				params.Amount,
			)
		default:
			return TransferERC20Token(
				params.Nonce,
				params.GasPrice,
				params.GasLimit,
				common.HexToAddress(sourceToken.ContractAddress),
				params.DestinationAddress,
				params.Amount,
			)
		}
	}

	switch sourceToken.NativeToken {
	case TokenTypeNative:
		return BridgeNativeToken(
			params.Client,
			params.Nonce,
			params.GasPrice,
			params.GasLimit,
			common.HexToAddress(sourceToken.BridgeContract),
			uint32(destinationToken.NetworkId),
			params.DestinationAddress,
			uint32(tokenID),
			params.Amount,
		)
	default:
		return BridgeERC20Token(
			params.Client,
			params.Nonce,
			params.GasPrice,
			params.GasLimit,
			common.HexToAddress(sourceToken.BridgeContract),
			common.HexToAddress(destinationToken.ContractAddress),
			uint32(destinationToken.NetworkId),
			params.DestinationAddress,
			uint32(tokenID),
			params.Amount,
		)
	}
}

func TransferNativeToken(
	nonce uint64,
	gasPrice *big.Int,
	gasLimit uint64,
	destinationAddress common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	return types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &destinationAddress,
		Value:    amount,
		Data:     nil,
	}), nil
}

func TransferERC20Token(
	nonce uint64,
	gasPrice *big.Int,
	gasLimit uint64,
	tokenAddress common.Address,
	destinationAddress common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	contractABI, err := contracts.ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	inputData, err := contractABI.Pack(
		"transfer",
		destinationAddress,
		amount,
	)
	return types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &tokenAddress,
		Value:    big.NewInt(0),
		Data:     inputData,
	}), nil
}

func BridgeNativeToken(
	client *ethclient.Client,
	nonce uint64,
	gasPrice *big.Int,
	gasLimit uint64,
	bridgeContract common.Address,
	destinationNetwork uint32,
	destinationAddress common.Address,
	tokenId uint32,
	amount *big.Int,
) (*types.Transaction, error) {
	bridgeInstance, _ := contracts.NewBridge(bridgeContract, client)
	fee, _ := bridgeInstance.EstimateBridgeToken(
		&bind.CallOpts{},
		destinationNetwork,
		common.HexToAddress("0x0"),
		0,
		big.NewInt(0),
		nil,
		big.NewInt(0),
	)
	formatAmount, _ := bridgeInstance.FormatAmount(&bind.CallOpts{}, tokenId, amount)

	bridgeABI, err := contracts.BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	inputData, err := bridgeABI.Pack(
		"bridgeToken",
		destinationNetwork,
		destinationAddress,
		tokenId,
		formatAmount,
		[]byte{},
		fee.TotalFee,
	)
	if err != nil {
		return nil, err
	}
	return types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &bridgeContract,
		Value:    formatAmount.Add(formatAmount, fee.TotalFee),
		Data:     inputData,
	}), nil
}

func BridgeERC20Token(
	client *ethclient.Client,
	nonce uint64,
	gasPrice *big.Int,
	gasLimit uint64,
	bridgeContract common.Address,
	tokenAddress common.Address,
	destinationNetwork uint32,
	destinationAddress common.Address,
	tokenId uint32,
	amount *big.Int,
) (*types.Transaction, error) {
	//erc20ABI, err := contracts.ERC20MetaData.GetAbi()
	//if err != nil {
	//	return nil, err
	//}
	bridgeInstance, _ := contracts.NewBridge(bridgeContract, client)
	formatAmount, _ := bridgeInstance.FormatAmount(&bind.CallOpts{}, tokenId, amount)
	//callData, err := erc20ABI.Pack(
	//	"transfer",
	//	destinationAddress,
	//	amount,
	//)
	fee, _ := bridgeInstance.EstimateBridgeToken(
		&bind.CallOpts{},
		destinationNetwork,
		common.HexToAddress("0x0"),
		0,
		big.NewInt(0),
		[]byte{},
		big.NewInt(0),
	)
	bridgeABI, err := contracts.BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	inputData, err := bridgeABI.Pack(
		"bridgeToken",
		destinationNetwork,
		destinationAddress,
		tokenId,
		formatAmount,
		[]byte{},
		fee.TotalFee,
	)
	return types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &bridgeContract,
		Value:    fee.TotalFee,
		Data:     inputData,
	}), nil
}
