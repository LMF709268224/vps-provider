package blockchain

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/xerrors"
)

func SendMessage() error {
	address := common.HexToAddress("0x61c6E61e8a5DC4117f4Aa72Bfa077b8B166b4a83")
	privateKeyStr := "3c3633bfaa3f8cfc2df9169d763eda6a8fb06d632e553f969f9dd2edd64dd11b"
	url := "http://183.60.189.250:1451/rpc/v1"

	client, err := ethclient.Dial(url)
	if err != nil {
		return xerrors.Errorf("Dial err:%s", err.Error())
	}

	o, err := NewOrder(address, client)
	if err != nil {
		return xerrors.Errorf("NewOrder err:%s", err.Error())
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return xerrors.Errorf("HexToECDSA err:%s", err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return xerrors.Errorf("publicKey err:%s", err.Error())
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31415926))
	if err != nil {
		return xerrors.Errorf("NewKeyedTransactorWithChainID err:%s", err.Error())
	}
	auth.GasLimit = uint64(3000000)
	auth.From = fromAddress

	tr, err := o.PlaceOrder(auth, big.NewInt(100), big.NewInt(10235))
	if err != nil {
		return xerrors.Errorf("PlaceOrder err:%s", err.Error())
	}

	fmt.Println(tr)
	return nil
}

// [第三方支付ID]VPSOrder
// []

type VPSOrder struct {
	Amount int64  `json:"amount"`
	To     string `json:"to"`
}
