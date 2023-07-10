package blockchain

import (
	"context"
	"encoding/json"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/lotus/chain/types"
	logging "github.com/ipfs/go-log/v2"
	"github.com/myxtype/filecoin-client"
	localfilecointypes "github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
)

var log = logging.Logger("blockchain")

const (
	lotusAddress = "http://183.60.189.250:1451/rpc/v1"
	authToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.3cIC5RdwN2xXPgBY6fkJGUkfVeCAnZVo3zVUls1oDes"
)

// MpoolPush 提交消息
func MpoolPush(body []byte) (string, error) {
	s := &localfilecointypes.SignedMessage{}
	err := json.Unmarshal(body, s)
	if err != nil {
		log.Errorf("MpoolPush Unmarshal error: %v", err)
		return "", err
	}

	// log.Infoln("url : ", url)
	client := filecoin.NewClient(lotusAddress, authToken)

	mid, err := client.MpoolPush(context.Background(), s)
	if err != nil {
		log.Errorf("MpoolPush ---- err : %v", err)
		return "", err
	}
	log.Infof("MpoolPush success cid : %v", mid)
	return mid.String(), nil
}

// ParseFIL Parse FIL
func ParseFIL(value string) (big.Int, error) {
	f, err := types.ParseFIL(value)
	if err != nil {
		return big.Int{}, err
	}

	amt := abi.TokenAmount(f)

	return amt, nil
}

// MakeMessage 构建message
func MakeMessage(info Message, nonce uint64) (*localfilecointypes.Message, uint64, error) {
	from, err := address.NewFromString(info.MFrom)
	if err != nil {
		return nil, 0, err
	}

	to, err := address.NewFromString(info.MTo)
	if err != nil {
		return nil, 0, err
	}

	value := info.Value + "afil"
	val, err := ParseFIL(value)
	if err != nil {
		return nil, 0, err
	}

	client := filecoin.NewClient(lotusAddress, authToken)

	// 获取nonce
	if nonce == 0 {
		nonce, err = client.MpoolGetNonce(context.Background(), from)
		if err != nil {
			// log.Printf("MpoolPush MpoolGetNonce err : %v", err)
			a, err := client.StateGetActor(context.Background(), from, nil)
			if err != nil {
				log.Errorf("MpoolPush StateGetActor err :%v\n", err)
				return nil, 0, err
			}
			nonce = a.Nonce + 1
		}
	}
	log.Infof("MpoolPush val:%v,nonce:%v", val, nonce)

	msg := &localfilecointypes.Message{
		Version:    0,
		To:         to,
		From:       from,
		Nonce:      nonce,
		Value:      val,
		GasLimit:   0,
		GasFeeCap:  abi.NewTokenAmount(0),
		GasPremium: abi.NewTokenAmount(0),
		Method:     uint64(builtin.MethodSend),
		Params:     nil,
	}
	// 默认 最大手续费0.0001 FIL 100000 nanoF 100000000000000 aF
	maxFee := filecoin.FromFil(decimal.NewFromFloat(0.0001))

	if info.GasLimit != "" && info.GasLimit != "0" {
		// 可接受手续费上限 单位:nanoFIL
		maxNano, err := decimal.NewFromString(info.GasLimit)
		if err != nil {
			log.Errorf("GasLimit to decinak err : %v", err)
			return nil, 0, err
		}

		// 手续费上限 从nanoFIL 转成 FIL
		maxFIL := maxNano.Div(decimal.NewFromFloat(float64(1000000000)))
		// 手续费上限 从 FIL 转成 aFIL
		maxFee = filecoin.FromFil(maxFIL)
	}
	// log.Infof("maxNano:%s,\nmaxFIL:%v,\nmaxFee:%v\n", maxNano, maxFIL, maxFee)

	// 估算GasLimit
	msg, err = client.GasEstimateMessageGas(context.Background(), msg, &localfilecointypes.MessageSendSpec{MaxFee: maxFee}, nil)
	if err != nil {
		log.Errorf("from:%s \nto:%s\n , MpoolPush GasEstimateMessageGas err : %v", from, to, err)
		return nil, 0, err
	}

	gas := msg.GasLimit * msg.GasFeeCap.Int64()
	gnano := decimal.NewFromFloat(float64(gas)).Div(decimal.NewFromFloat(float64(1000000000)))
	log.Infof("估算gas :%v nanoFIL", gnano)

	return msg, nonce, nil
}

// Message 待确认交易表
type Message struct {
	TransID  int64  `json:"column:trans_id"` // 订单ID
	ID       int64  `gorm:"column:id"`
	MTo      string `gorm:"column:mto"`      // to
	MFrom    string `gorm:"column:mfrom"`    // from
	Value    string `gorm:"column:value"`    // 交易值
	GasLimit string `gorm:"column:gaslimit"` // 可接受手续费上限 单位:nanoFIL
	PType    string `gorm:"column:ptype"`    // 类型
	Nonce    uint64 `gorm:"column:nonce"`    // 交易计数
}
