/*
Copyright 2022 NVIDIA CORPORATION & AFFILIATES
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	upgrade "github.com/NVIDIA/operator-libs/pkg/upgrade"
	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/api/core/v1"
)

// PodManager is an autogenerated mock type for the PodManager type
type PodManager struct {
	mock.Mock
}

// ScheduleCheckOnPodCompletion provides a mock function with given fields: _a0, _a1
func (_m *PodManager) ScheduleCheckOnPodCompletion(_a0 context.Context, _a1 *upgrade.PodManagerConfig) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *upgrade.PodManagerConfig) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SchedulePodEviction provides a mock function with given fields: _a0, _a1
func (_m *PodManager) SchedulePodEviction(_a0 context.Context, _a1 *upgrade.PodManagerConfig) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *upgrade.PodManagerConfig) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SchedulePodsRestart provides a mock function with given fields: _a0, _a1
func (_m *PodManager) SchedulePodsRestart(_a0 context.Context, _a1 []*v1.Pod) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*v1.Pod) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPodManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewPodManager creates a new instance of PodManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPodManager(t mockConstructorTestingTNewPodManager) *PodManager {
	mock := &PodManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
