// This file is part of SGU Go Client.
//
// (c) SGU Ecosystem <info@SGU.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"context"
	"fmt"
	"net/http"
)

// BlocksService handles communication with the blocks related
// methods of the SGU Core API - Version 2.
type BlocksService Service

type PaginationHeight struct {
	Page    int    `url:"page"`
	Limit   int    `url:"limit"`
	Height  int    `url:"height"`
	//BlockId string `url:"blockId"`
}



// Get all blocks.
func (s *BlocksService) List(ctx context.Context, query *Pagination) (*Blocks, *http.Response, error) {
	var responseStruct *Blocks
	resp, err := s.client.SendRequest(ctx, "GET", "blocks", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}


	return responseStruct, resp, err
}


func (s *BlocksService) ListByHeight(ctx context.Context, query *PaginationHeight) (*Blocks, *http.Response, error) {
	var responseStruct *Blocks
	resp, err := s.client.SendRequest(ctx, "GET", "blocks", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}


	return responseStruct, resp, err
}

// Get a block by the given id.
func (s *BlocksService) Get(ctx context.Context, id int) (*GetBlock, *http.Response, error) {
	uri := fmt.Sprintf("blocks/%v", id)

	var responseStruct *GetBlock
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions by the given block.
func (s *BlocksService) Transactions(ctx context.Context, id int, query *Pagination) (*GetBlockTransactions, *http.Response, error) {
	uri := fmt.Sprintf("blocks/%v/transactions", id)

	var responseStruct *GetBlockTransactions
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all blocks by the given criteria.
func (s *BlocksService) Search(ctx context.Context, query *Pagination, body *BlocksSearchRequest) (*Blocks, *http.Response, error) {
	var responseStruct *Blocks
	resp, err := s.client.SendRequest(ctx, "POST", "blocks/search", query, body, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}
