package transactions

import (
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
)

const (
	TokenInfoURL = "https://api.goose.farm/v5/common/token/info?token_symbol="
)

type TokenResponse struct {
	Code    int       `json:"code"`
	Data    TokenInfo `json:"data"`
	Message string    `json:"message"`
}

type TokenInfo struct {
	TokenId     int32          `json:"token_id"`
	TokenSymbol string         `json:"token_symbol"`
	TokenIcon   string         `json:"token_icon"`
	ChainList   []*TokenDetail `json:"chain_list"`
}

type TokenDetail struct {
	TokenName       string `json:"token_name"`
	NativeToken     int32  `json:"native_token"`
	ChainId         int64  `json:"chain_id"`
	NetworkId       int64  `json:"network_id"`
	BridgeContract  string `json:"bridge_contract"`
	ChainSymbol     string `json:"chain_symbol"`
	ChainIcon       string `json:"chain_icon"`
	ContractAddress string `json:"contract_address"`
	SharedDecimals  uint32 `json:"shared_decimals"`
	LocalDecimals   uint32 `json:"local_decimals"`
}

func queryTokenInfo(tokenSymbol string, sourceChainID, destinationChainID int64) (int32, *TokenDetail, *TokenDetail, error) {
	resp, err := http.Get(TokenInfoURL + tokenSymbol)
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, nil, err
	}
	var data TokenResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, nil, nil, err
	}
	if data.Code != 0 {
		return 0, nil, nil, fmt.Errorf("query failed:%s", data.Message)
	}

	var sourceToken, destinationToken *TokenDetail
	for _, v := range data.Data.ChainList {
		if v.ChainId == sourceChainID {
			sourceToken = v
		}
		if v.ChainId == destinationChainID {
			destinationToken = v
		}
	}
	return data.Data.TokenId, sourceToken, destinationToken, nil
}
