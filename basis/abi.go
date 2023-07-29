package basis

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"golang.org/x/xerrors"
)

func main() {
	err := testWatch()
	fmt.Println("\nSendMessage err:", err)
}

func testWatch() error {
	client, err := ethclient.Dial("wss://183.60.189.250:12340/rpc/v0")
	if err != nil {
		return xerrors.Errorf("Dial err:%s", err.Error())
	}

	headers := make(chan *types.Header, 500)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		return xerrors.Errorf("SubscribeNewHead err:%s", err.Error())
	}

	fmt.Printf("tx sent: %s \n", sub)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("------------------", header.Hash().Hex())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			for _, transaction := range block.Transactions() {
				fmt.Println("hash:", transaction.Hash(), " Type:", transaction.Type(), " Data:", string(transaction.Data()))
			}
		}
	}
}

func testTransfer() error {
	client, err := ethclient.Dial("https://api.calibration.node.glif.io/rpc/v1")
	if err != nil {
		return xerrors.Errorf("Dial err:%s", err.Error())
	}

	privateKey, err := crypto.HexToECDSA("3c3633bfaa3f8cfc2df9169d763eda6a8fb06d632e553f969f9dd2edd64dd11b")
	if err != nil {
		return xerrors.Errorf("HexToECDSA err:%s", err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return xerrors.New("publicKey err:")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress("0x5feaAc40B8eB3575794518bC0761cB4A95838ccF")
	tokenAddress := common.HexToAddress("0x54677876506eB9ae9095ca7438cCD347b3159660")
	transferFnSignature := []byte("transfer(address,uint256)")

	myAbi, err := abi.NewAbi(tokenAddress, client)
	if err != nil {
		return xerrors.Errorf("NewAbi err:%s", err.Error())
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("100000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		return xerrors.Errorf("EstimateGas err:%s", err.Error())
	}

	fmt.Println(gasLimit) // 23256

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return xerrors.Errorf("NetworkID err:%s", err.Error())
	}

	signer := types.LatestSignerForChainID(chainID)
	to := &bind.TransactOpts{
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(transaction, signer, privateKey)
		},
		From:     fromAddress,
		Context:  context.Background(),
		GasLimit: gasLimit,
	}

	signedTx, err := myAbi.Transfer(to, toAddress, amount)
	if err != nil {
		return xerrors.Errorf("Transfer err:%s", err.Error())
	}

	fmt.Printf("tx sent: %s \n", signedTx.Hash().Hex())
	return nil
}
