/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */
package suchain

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/blocktree/openwallet/log"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)


type Client struct {
	BaseURL   string
	Debug     bool
	ErrorTime int
	lock      sync.Mutex
	DelayTime int64
}

type Response struct {
	Id      int         `json:"id"`
	Version string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
}


//func (this *Client) PushTransaction(packedTx interface{}) (*crypto.Transaction, error) {
//	params := []interface{}{
//		//appendOxToAddress(addr),
//		"network_broadcast_api",
//		"broadcast_transaction_synchronous",
//		[]interface{}{packedTx},
//	}
//	result, err := this.Call("call", 1, params)
//	if err != nil {
//		log.Errorf("pushTransaction faield, err = %v \n", err)
//		return nil, err
//	}
//
//	if result.Type != gjson.JSON {
//		log.Errorf("result of pushTransaction type error")
//		return nil, errors.New("result of pushTransaction type error")
//	}
//
//	var apiTransResult *ApiTransResult
//	err = json.Unmarshal([]byte(result.Raw), &apiTransResult)
//	if err != nil {
//		log.Errorf("pushTransaction decode json [%v] failed, err=%v", []byte(result.Raw), err)
//		return nil, err
//	}
//
//	return apiTransResult, nil
//}

func (c *Client) Call(method string, id int64, params []interface{}) (*gjson.Result, error) {
	time.Sleep(time.Duration(c.DelayTime) * time.Millisecond)
	c.lock.Lock()
	defer c.lock.Unlock()
	authHeader := req.Header{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	body := make(map[string]interface{}, 0)
	body["jsonrpc"] = "2.0"
	body["id"] = id
	body["method"] = method
	body["params"] = params

	if c.Debug {
		log.Debug("Start Request API...")
	}

	r, err := req.Post(c.BaseURL, req.BodyJSON(&body), authHeader)

	if c.Debug {
		log.Debug("Request API Completed")
	}

	if c.Debug {
		log.Debugf("%+v\n", r)
	}

	if err != nil {
		return nil, err
	}

	resp := gjson.ParseBytes(r.Bytes())
	err = isError(&resp)
	if err != nil {
		return nil, err
	}

	result := resp.Get("result")

	return &result, nil
}

//isError 是否报错
func isError(result *gjson.Result) error {
	var (
		err error
	)

	if !result.Get("error").IsObject() {

		if !result.Get("result").Exists() {
			return errors.New("Response is empty! ")
		}

		return nil
	}

	errInfo := fmt.Sprintf("[%d]%s",
		result.Get("error.code").Int(),
		result.Get("error.message").String())
	err = errors.New(errInfo)

	return err
}
