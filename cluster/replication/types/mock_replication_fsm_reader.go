//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by mockery v2.53.2. DO NOT EDIT.

package types

import mock "github.com/stretchr/testify/mock"

// MockReplicationFSMReader is an autogenerated mock type for the ReplicationFSMReader type
type MockReplicationFSMReader struct {
	mock.Mock
}

type MockReplicationFSMReader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReplicationFSMReader) EXPECT() *MockReplicationFSMReader_Expecter {
	return &MockReplicationFSMReader_Expecter{mock: &_m.Mock}
}

// FilterOneShardReplicasReadWrite provides a mock function with given fields: collection, shard, shardReplicasLocation
func (_m *MockReplicationFSMReader) FilterOneShardReplicasReadWrite(collection string, shard string, shardReplicasLocation []string) ([]string, []string) {
	ret := _m.Called(collection, shard, shardReplicasLocation)

	if len(ret) == 0 {
		panic("no return value specified for FilterOneShardReplicasReadWrite")
	}

	var r0 []string
	var r1 []string
	if rf, ok := ret.Get(0).(func(string, string, []string) ([]string, []string)); ok {
		return rf(collection, shard, shardReplicasLocation)
	}
	if rf, ok := ret.Get(0).(func(string, string, []string) []string); ok {
		r0 = rf(collection, shard, shardReplicasLocation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, []string) []string); ok {
		r1 = rf(collection, shard, shardReplicasLocation)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	return r0, r1
}

// MockReplicationFSMReader_FilterOneShardReplicasReadWrite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FilterOneShardReplicasReadWrite'
type MockReplicationFSMReader_FilterOneShardReplicasReadWrite_Call struct {
	*mock.Call
}

// FilterOneShardReplicasReadWrite is a helper method to define mock.On call
//   - collection string
//   - shard string
//   - shardReplicasLocation []string
func (_e *MockReplicationFSMReader_Expecter) FilterOneShardReplicasReadWrite(collection interface{}, shard interface{}, shardReplicasLocation interface{}) *MockReplicationFSMReader_FilterOneShardReplicasReadWrite_Call {
	return &MockReplicationFSMReader_FilterOneShardReplicasReadWrite_Call{Call: _e.mock.On("FilterOneShardReplicasReadWrite", collection, shard, shardReplicasLocation)}
}

func (_c *MockReplicationFSMReader_FilterOneShardReplicasReadWrite_Call) Run(run func(collection string, shard string, shardReplicasLocation []string)) *MockReplicationFSMReader_FilterOneShardReplicasReadWrite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *MockReplicationFSMReader_FilterOneShardReplicasReadWrite_Call) Return(_a0 []string, _a1 []string) *MockReplicationFSMReader_FilterOneShardReplicasReadWrite_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReplicationFSMReader_FilterOneShardReplicasReadWrite_Call) RunAndReturn(run func(string, string, []string) ([]string, []string)) *MockReplicationFSMReader_FilterOneShardReplicasReadWrite_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockReplicationFSMReader creates a new instance of MockReplicationFSMReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReplicationFSMReader(t interface {
	mock.TestingT
	Cleanup(func())
},
) *MockReplicationFSMReader {
	mock := &MockReplicationFSMReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
