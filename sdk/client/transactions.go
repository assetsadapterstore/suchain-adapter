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

// TransactionsService handles communication with the transactions related
// methods of the SGU Core API - Version 2.
type TransactionsService Service

type PaginationBlock struct {
	Page    int    `url:"page"`
	Limit   int    `url:"limit"`
	//Height  int    `url:"height"`
	BlockId string `url:"blockId"`
}


// Get all transactions.
func (s *TransactionsService) List(ctx context.Context, query *Pagination) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", "transactions", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions.
func (s *TransactionsService) ListByBlockId(ctx context.Context, query *PaginationBlock) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	query.Page = 1
	resp, err := s.client.SendRequest(ctx, "GET", "transactions", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions.
func (s *TransactionsService) ListById(ctx context.Context, query *Pagination,id string) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions

	uri := fmt.Sprintf("transactions/%v", id)

	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}



// Create a new transaction.
func (s *TransactionsService) Create(ctx context.Context, body *CreateTransactionRequest) (*CreateTransaction, *http.Response, error) {
	var responseStruct *CreateTransaction
	resp, err := s.client.SendRequest(ctx, "POST", "transactions", nil, body, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a transaction by the given id.
func (s *TransactionsService) Get(ctx context.Context, id string) (*GetTransaction, *http.Response, error) {
	uri := fmt.Sprintf("transactions/%v", id)

	var responseStruct *GetTransaction
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all unconfirmed transactions.
func (s *TransactionsService) ListUnconfirmed(ctx context.Context, query *Pagination) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", "transactions/unconfirmed", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get an unconfirmed transaction by the given id.
func (s *TransactionsService) GetUnconfirmed(ctx context.Context, id string) (*GetTransaction, *http.Response, error) {
	uri := fmt.Sprintf("transactions/unconfirmed/%v", id)

	var responseStruct *GetTransaction
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all transactions by the given criteria.
func (s *TransactionsService) Search(ctx context.Context, query *Pagination, body *TransactionsSearchRequest) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "POST", "transactions/search", query, body, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a list of valid transaction types.
func (s *TransactionsService) Types(ctx context.Context) (*TransactionTypes, *http.Response, error) {
	var responseStruct *TransactionTypes
	resp, err := s.client.SendRequest(ctx, "GET", "transactions/types", nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}
