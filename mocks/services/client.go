// Code generated by mockery 2.9.4. DO NOT EDIT.

package services

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/coinbase/rosetta-sdk-go/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Balance provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Client) Balance(_a0 context.Context, _a1 *types.AccountIdentifier, _a2 *types.PartialBlockIdentifier, _a3 []*types.Currency) (*types.AccountBalanceResponse, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.AccountBalanceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountIdentifier, *types.PartialBlockIdentifier, []*types.Currency) *types.AccountBalanceResponse); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.AccountBalanceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.AccountIdentifier, *types.PartialBlockIdentifier, []*types.Currency) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Block provides a mock function with given fields: _a0, _a1
func (_m *Client) Block(_a0 context.Context, _a1 *types.PartialBlockIdentifier) (*types.Block, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, *types.PartialBlockIdentifier) *types.Block); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.PartialBlockIdentifier) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Call provides a mock function with given fields: ctx, request
func (_m *Client) Call(ctx context.Context, request *types.CallRequest) (*types.CallResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *types.CallResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.CallRequest) *types.CallResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.CallResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.CallRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateGas provides a mock function with given fields: ctx, msg
func (_m *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	ret := _m.Called(ctx, msg)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg) uint64); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingNonceAt provides a mock function with given fields: _a0, _a1
func (_m *Client) PendingNonceAt(_a0 context.Context, _a1 common.Address) (uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransaction provides a mock function with given fields: ctx, tx
func (_m *Client) SendTransaction(ctx context.Context, tx *coretypes.Transaction) error {
	ret := _m.Called(ctx, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *coretypes.Transaction) error); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with given fields: _a0
func (_m *Client) Status(_a0 context.Context) (*types.BlockIdentifier, int64, *types.SyncStatus, []*types.Peer, error) {
	ret := _m.Called(_a0)

	var r0 *types.BlockIdentifier
	if rf, ok := ret.Get(0).(func(context.Context) *types.BlockIdentifier); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockIdentifier)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context) int64); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 *types.SyncStatus
	if rf, ok := ret.Get(2).(func(context.Context) *types.SyncStatus); ok {
		r2 = rf(_a0)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*types.SyncStatus)
		}
	}

	var r3 []*types.Peer
	if rf, ok := ret.Get(3).(func(context.Context) []*types.Peer); ok {
		r3 = rf(_a0)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).([]*types.Peer)
		}
	}

	var r4 error
	if rf, ok := ret.Get(4).(func(context.Context) error); ok {
		r4 = rf(_a0)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// SuggestGasPrice provides a mock function with given fields: ctx, gasPrice
func (_m *Client) SuggestGasPrice(ctx context.Context, gasPrice *big.Int) (*big.Int, error) {
	ret := _m.Called(ctx, gasPrice)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *big.Int); ok {
		r0 = rf(ctx, gasPrice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, gasPrice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
