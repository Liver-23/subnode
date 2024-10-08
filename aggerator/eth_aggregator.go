package aggerator

import (
	"encoding/json"
	"github.com/notional-labs/subnode/state"
	"github.com/notional-labs/subnode/utils"
	"log"
	"net/http"
)

func Eth_BathRequest(w http.ResponseWriter, jsonBody []byte) {
	log.Printf("Eth_BathRequest...")
	var arr []json.RawMessage

	err := json.Unmarshal(jsonBody, &arr)
	if err != nil {
		_ = utils.SendError(w)
		return
	}

	var arrSize = len(arr)
	var arrRes = make([]json.RawMessage, arrSize)

	for i, s := range arr {
		bodyChild, err := utils.FetchJsonRpcOverHttp("http://localhost:39545", s)
		if err != nil {
			_ = utils.SendError(w)
			return
		}

		arrRes[i] = bodyChild
	}

	jsonBytes, err := json.Marshal(arrRes)
	if err != nil {
		_ = utils.SendError(w)
		return
	}

	_ = utils.SendResult(w, jsonBytes)
}

func Eth_getBlockTransactionCountByHash(w http.ResponseWriter, jsonBody []byte) {
	for i, s := range state.PoolEth {
		ethUrl := s.Backend.Eth
		body, err := utils.FetchJsonRpcOverHttp(ethUrl, jsonBody)
		if err != nil {
			continue
		}

		var j0 interface{}
		err = json.Unmarshal(body, &j0)
		if err != nil {
			_ = utils.SendError(w)
			return
		}

		if m0, ok := j0.(map[string]interface{}); ok {
			if (m0["result"] != nil) || (i >= len(state.PoolEth)-1) { // found result or last node, send it
				_ = utils.SendResult(w, body)
				return
			}
		}

	}

	_ = utils.SendError(w)
}

func Eth_getBlockByHash(w http.ResponseWriter, jsonBody []byte) {
	//log.Println("Eth_getBlockByHash")
	for i, s := range state.PoolEth {
		ethUrl := s.Backend.Eth
		//log.Println("Eth_getBlockByHash ethUrl=", ethUrl)
		body, err := utils.FetchJsonRpcOverHttp(ethUrl, jsonBody)
		if err != nil {
			//log.Println("Eth_getBlockByHash Error, continue", err)
			continue
		}

		var j0 interface{}
		err = json.Unmarshal(body, &j0)
		if err != nil {
			//log.Println("Eth_getBlockByHash SendError", err)
			_ = utils.SendError(w)
			return
		}

		if m0, ok := j0.(map[string]interface{}); ok {
			if (m0["result"] != nil) || (i >= len(state.PoolEth)-1) { // found result or last node, send it
				//log.Println("Eth_getBlockByHash SendResult")
				_ = utils.SendResult(w, body)
				return
			}
		}

	}

	_ = utils.SendError(w)
}

func Eth_getTransactionByHash(w http.ResponseWriter, jsonBody []byte) {
	//log.Println("Eth_getTransactionByHash")
	for i, s := range state.PoolEth {
		ethUrl := s.Backend.Eth
		//log.Println("Eth_getTransactionByHash ethUrl=", ethUrl)
		body, err := utils.FetchJsonRpcOverHttp(ethUrl, jsonBody)
		if err != nil {
			continue
		}

		var j0 interface{}
		err = json.Unmarshal(body, &j0)
		if err != nil {
			//log.Println("Eth_getTransactionByHash SendError")
			_ = utils.SendError(w)
			return
		}

		if m0, ok := j0.(map[string]interface{}); ok {
			if (m0["result"] != nil) || (i >= len(state.PoolEth)-1) { // found result or last node, send it
				//log.Println("Eth_getTransactionByHash SendResult")
				_ = utils.SendResult(w, body)
				return
			}
		}

	}

	_ = utils.SendError(w)
}

func Eth_getTransactionByBlockHashAndIndex(w http.ResponseWriter, jsonBody []byte) {
	for i, s := range state.PoolEth {
		ethUrl := s.Backend.Eth
		body, err := utils.FetchJsonRpcOverHttp(ethUrl, jsonBody)
		if err != nil {
			continue
		}

		var j0 interface{}
		err = json.Unmarshal(body, &j0)
		if err != nil {
			_ = utils.SendError(w)
			return
		}

		if m0, ok := j0.(map[string]interface{}); ok {
			if (m0["result"] != nil) || (i >= len(state.PoolEth)-1) { // found result or last node, send it
				_ = utils.SendResult(w, body)
				return
			}
		}

	}

	_ = utils.SendError(w)
}

func Eth_getTransactionReceipt(w http.ResponseWriter, jsonBody []byte) {
	for i, s := range state.PoolEth {
		ethUrl := s.Backend.Eth
		body, err := utils.FetchJsonRpcOverHttp(ethUrl, jsonBody)
		if err != nil {
			continue
		}

		var j0 interface{}
		err = json.Unmarshal(body, &j0)
		if err != nil {
			_ = utils.SendError(w)
			return
		}

		if m0, ok := j0.(map[string]interface{}); ok {
			if (m0["result"] != nil) || (i >= len(state.PoolEth)-1) { // found result or last node, send it
				_ = utils.SendResult(w, body)
				return
			}
		}

	}

	_ = utils.SendError(w)
}
