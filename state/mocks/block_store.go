// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/providenetwork/tendermint/types"
)

// BlockStore is an autogenerated mock type for the BlockStore type
type BlockStore struct {
	mock.Mock
}

// Base provides a mock function with given fields:
func (_m *BlockStore) Base() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Height provides a mock function with given fields:
func (_m *BlockStore) Height() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// LoadBaseMeta provides a mock function with given fields:
func (_m *BlockStore) LoadBaseMeta() *types.BlockMeta {
	ret := _m.Called()

	var r0 *types.BlockMeta
	if rf, ok := ret.Get(0).(func() *types.BlockMeta); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockMeta)
		}
	}

	return r0
}

// LoadBlock provides a mock function with given fields: height
func (_m *BlockStore) LoadBlock(height int64) *types.Block {
	ret := _m.Called(height)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(int64) *types.Block); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	return r0
}

// LoadBlockByHash provides a mock function with given fields: hash
func (_m *BlockStore) LoadBlockByHash(hash []byte) *types.Block {
	ret := _m.Called(hash)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func([]byte) *types.Block); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	return r0
}

// LoadBlockCommit provides a mock function with given fields: height
func (_m *BlockStore) LoadBlockCommit(height int64) *types.Commit {
	ret := _m.Called(height)

	var r0 *types.Commit
	if rf, ok := ret.Get(0).(func(int64) *types.Commit); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Commit)
		}
	}

	return r0
}

// LoadBlockMeta provides a mock function with given fields: height
func (_m *BlockStore) LoadBlockMeta(height int64) *types.BlockMeta {
	ret := _m.Called(height)

	var r0 *types.BlockMeta
	if rf, ok := ret.Get(0).(func(int64) *types.BlockMeta); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockMeta)
		}
	}

	return r0
}

// LoadBlockPart provides a mock function with given fields: height, index
func (_m *BlockStore) LoadBlockPart(height int64, index int) *types.Part {
	ret := _m.Called(height, index)

	var r0 *types.Part
	if rf, ok := ret.Get(0).(func(int64, int) *types.Part); ok {
		r0 = rf(height, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Part)
		}
	}

	return r0
}

// LoadSeenCommit provides a mock function with given fields:
func (_m *BlockStore) LoadSeenCommit() *types.Commit {
	ret := _m.Called()

	var r0 *types.Commit
	if rf, ok := ret.Get(0).(func() *types.Commit); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Commit)
		}
	}

	return r0
}

// PruneBlocks provides a mock function with given fields: height
func (_m *BlockStore) PruneBlocks(height int64) (uint64, error) {
	ret := _m.Called(height)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(int64) uint64); ok {
		r0 = rf(height)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveBlock provides a mock function with given fields: block, blockParts, seenCommit
func (_m *BlockStore) SaveBlock(block *types.Block, blockParts *types.PartSet, seenCommit *types.Commit) {
	_m.Called(block, blockParts, seenCommit)
}

// Size provides a mock function with given fields:
func (_m *BlockStore) Size() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}
