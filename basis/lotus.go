package basis

import (
	"context"
	"fmt"
	"net/http"

	jsonrpc "github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

const (
	lotusAddress = "api.calibration.node.glif.io"
	authToken    = ""
)

// GetLotusHTTPAPI 获取lotus api
func GetLotusHTTPAPI() (lotusapi.FullNodeStruct, jsonrpc.ClientCloser, error) {
	headers := http.Header{"Authorization": []string{"Bearer " + authToken}}

	url := "http://" + lotusAddress + "/rpc/v0"

	var api lotusapi.FullNodeStruct
	closer, err := jsonrpc.NewMergeClient(context.Background(), url, "Filecoin",
		[]interface{}{&api.Internal, &api.CommonStruct.Internal}, headers)
	if err != nil {
		return api, nil, err
	}
	// log.Infoln("url : ", url)

	return api, closer, nil
}

// ChainGetTipSetByHeight 根据高度获取TipSet
func ChainGetTipSetByHeight(lotusAPI *lotusapi.FullNodeStruct, height int64) (*types.TipSet, error) {
	if lotusAPI == nil {
		lpai, closer, err := GetLotusHTTPAPI()
		if err != nil {
			return nil, err
		}
		defer closer()

		lotusAPI = &lpai
	}

	return lotusAPI.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(height), types.NewTipSetKey())
}

// ChainGetBlockMessages 获取区块消息
func ChainGetBlockMessages(lotusAPI *lotusapi.FullNodeStruct, mCid cid.Cid) (*lotusapi.BlockMessages, error) {
	if lotusAPI == nil {
		lpai, closer, err := GetLotusHTTPAPI()
		if err != nil {
			return nil, err
		}
		defer closer()

		lotusAPI = &lpai
	}

	return lotusAPI.ChainGetBlockMessages(context.Background(), mCid)
}

func workerHander(blocknumber int64) error {
	lotusAPI, closer, err := GetLotusHTTPAPI()
	if err != nil {
		return err
	}
	defer func() {
		closer()
		if r := recover(); r != nil {
			fmt.Errorf("捕获到的错误:%s\n", r)
		}
	}()

	// 获取 指定高度的 tipset
	tipSet, err := ChainGetTipSetByHeight(&lotusAPI, blocknumber)
	if err != nil {
		fmt.Errorf("%d workerHander ChainGetTipSetByHeight err : %v", blocknumber, err)
		return err
	}

	// log.Infof("%d 扫到块 Height : %v", num, tipSet.Height())
	msgMap := make(map[string]types.Message)
	blockhashMap := make(map[string]cid.Cid)

	for _, tscid := range tipSet.Cids() {
		cid := tscid

		info, err := ChainGetBlockMessages(&lotusAPI, cid)
		if err != nil {
			fmt.Errorf("workerHander ChainGetBlockMessages err : %s", err.Error())
			return err
		}

		for _, msg := range info.BlsMessages {
			msgMap[msg.Cid().String()] = *msg
			blockhashMap[msg.Cid().String()] = cid
		}

		for _, msg := range info.SecpkMessages {
			msgMap[msg.Cid().String()] = msg.Message
			blockhashMap[msg.Cid().String()] = cid
		}
	}

	messageHandle(msgMap, int64(tipSet.Height()), blockhashMap)

	return nil
}

// 检查地址
func messageHandle(mMap map[string]types.Message, height int64, bMap map[string]cid.Cid) {
	for mCid, msg := range mMap {

		// blockCid := bMap[mCid]

		fromAddress := msg.From.String()
		toAddress := msg.To.String()

		fmt.Printf("cid:%s,from:%s,to:%s \n", mCid, &fromAddress, toAddress)
	}
}
