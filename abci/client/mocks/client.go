// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	abcicli "github.com/providenetwork/tendermint/abci/client"

	log "github.com/providenetwork/tendermint/libs/log"

	mock "github.com/stretchr/testify/mock"

	types "github.com/providenetwork/tendermint/abci/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// ApplySnapshotChunkAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) ApplySnapshotChunkAsync(_a0 context.Context, _a1 types.RequestApplySnapshotChunk) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestApplySnapshotChunk) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestApplySnapshotChunk) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplySnapshotChunkSync provides a mock function with given fields: _a0, _a1
func (_m *Client) ApplySnapshotChunkSync(_a0 context.Context, _a1 types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseApplySnapshotChunk
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestApplySnapshotChunk) *types.ResponseApplySnapshotChunk); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseApplySnapshotChunk)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestApplySnapshotChunk) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeginBlockAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) BeginBlockAsync(_a0 context.Context, _a1 types.RequestBeginBlock) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestBeginBlock) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestBeginBlock) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeginBlockSync provides a mock function with given fields: _a0, _a1
func (_m *Client) BeginBlockSync(_a0 context.Context, _a1 types.RequestBeginBlock) (*types.ResponseBeginBlock, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseBeginBlock
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestBeginBlock) *types.ResponseBeginBlock); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseBeginBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestBeginBlock) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckTxAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) CheckTxAsync(_a0 context.Context, _a1 types.RequestCheckTx) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestCheckTx) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestCheckTx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckTxSync provides a mock function with given fields: _a0, _a1
func (_m *Client) CheckTxSync(_a0 context.Context, _a1 types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseCheckTx
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestCheckTx) *types.ResponseCheckTx); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseCheckTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestCheckTx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CommitAsync provides a mock function with given fields: _a0
func (_m *Client) CommitAsync(_a0 context.Context) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context) *abcicli.ReqRes); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CommitSync provides a mock function with given fields: _a0
func (_m *Client) CommitSync(_a0 context.Context) (*types.ResponseCommit, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseCommit
	if rf, ok := ret.Get(0).(func(context.Context) *types.ResponseCommit); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseCommit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeliverTxAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) DeliverTxAsync(_a0 context.Context, _a1 types.RequestDeliverTx) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestDeliverTx) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestDeliverTx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeliverTxSync provides a mock function with given fields: _a0, _a1
func (_m *Client) DeliverTxSync(_a0 context.Context, _a1 types.RequestDeliverTx) (*types.ResponseDeliverTx, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseDeliverTx
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestDeliverTx) *types.ResponseDeliverTx); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseDeliverTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestDeliverTx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EchoAsync provides a mock function with given fields: ctx, msg
func (_m *Client) EchoAsync(ctx context.Context, msg string) (*abcicli.ReqRes, error) {
	ret := _m.Called(ctx, msg)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, string) *abcicli.ReqRes); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EchoSync provides a mock function with given fields: ctx, msg
func (_m *Client) EchoSync(ctx context.Context, msg string) (*types.ResponseEcho, error) {
	ret := _m.Called(ctx, msg)

	var r0 *types.ResponseEcho
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.ResponseEcho); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseEcho)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndBlockAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) EndBlockAsync(_a0 context.Context, _a1 types.RequestEndBlock) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestEndBlock) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestEndBlock) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndBlockSync provides a mock function with given fields: _a0, _a1
func (_m *Client) EndBlockSync(_a0 context.Context, _a1 types.RequestEndBlock) (*types.ResponseEndBlock, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseEndBlock
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestEndBlock) *types.ResponseEndBlock); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseEndBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestEndBlock) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Error provides a mock function with given fields:
func (_m *Client) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlushAsync provides a mock function with given fields: _a0
func (_m *Client) FlushAsync(_a0 context.Context) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context) *abcicli.ReqRes); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FlushSync provides a mock function with given fields: _a0
func (_m *Client) FlushSync(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InfoAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) InfoAsync(_a0 context.Context, _a1 types.RequestInfo) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestInfo) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestInfo) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InfoSync provides a mock function with given fields: _a0, _a1
func (_m *Client) InfoSync(_a0 context.Context, _a1 types.RequestInfo) (*types.ResponseInfo, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseInfo
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestInfo) *types.ResponseInfo); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestInfo) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitChainAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) InitChainAsync(_a0 context.Context, _a1 types.RequestInitChain) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestInitChain) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestInitChain) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitChainSync provides a mock function with given fields: _a0, _a1
func (_m *Client) InitChainSync(_a0 context.Context, _a1 types.RequestInitChain) (*types.ResponseInitChain, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseInitChain
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestInitChain) *types.ResponseInitChain); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseInitChain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestInitChain) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsRunning provides a mock function with given fields:
func (_m *Client) IsRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ListSnapshotsAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) ListSnapshotsAsync(_a0 context.Context, _a1 types.RequestListSnapshots) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestListSnapshots) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestListSnapshots) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSnapshotsSync provides a mock function with given fields: _a0, _a1
func (_m *Client) ListSnapshotsSync(_a0 context.Context, _a1 types.RequestListSnapshots) (*types.ResponseListSnapshots, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseListSnapshots
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestListSnapshots) *types.ResponseListSnapshots); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseListSnapshots)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestListSnapshots) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadSnapshotChunkAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) LoadSnapshotChunkAsync(_a0 context.Context, _a1 types.RequestLoadSnapshotChunk) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestLoadSnapshotChunk) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestLoadSnapshotChunk) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadSnapshotChunkSync provides a mock function with given fields: _a0, _a1
func (_m *Client) LoadSnapshotChunkSync(_a0 context.Context, _a1 types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseLoadSnapshotChunk
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestLoadSnapshotChunk) *types.ResponseLoadSnapshotChunk); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseLoadSnapshotChunk)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestLoadSnapshotChunk) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OfferSnapshotAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) OfferSnapshotAsync(_a0 context.Context, _a1 types.RequestOfferSnapshot) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestOfferSnapshot) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestOfferSnapshot) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OfferSnapshotSync provides a mock function with given fields: _a0, _a1
func (_m *Client) OfferSnapshotSync(_a0 context.Context, _a1 types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseOfferSnapshot
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestOfferSnapshot) *types.ResponseOfferSnapshot); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseOfferSnapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestOfferSnapshot) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnReset provides a mock function with given fields:
func (_m *Client) OnReset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnStart provides a mock function with given fields:
func (_m *Client) OnStart() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnStop provides a mock function with given fields:
func (_m *Client) OnStop() {
	_m.Called()
}

// QueryAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) QueryAsync(_a0 context.Context, _a1 types.RequestQuery) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestQuery) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuerySync provides a mock function with given fields: _a0, _a1
func (_m *Client) QuerySync(_a0 context.Context, _a1 types.RequestQuery) (*types.ResponseQuery, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseQuery
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestQuery) *types.ResponseQuery); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseQuery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Quit provides a mock function with given fields:
func (_m *Client) Quit() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Reset provides a mock function with given fields:
func (_m *Client) Reset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLogger provides a mock function with given fields: _a0
func (_m *Client) SetLogger(_a0 log.Logger) {
	_m.Called(_a0)
}

// SetResponseCallback provides a mock function with given fields: _a0
func (_m *Client) SetResponseCallback(_a0 abcicli.Callback) {
	_m.Called(_a0)
}

// Start provides a mock function with given fields:
func (_m *Client) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Client) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Client) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Wait provides a mock function with given fields:
func (_m *Client) Wait() {
	_m.Called()
}
