/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */
package engine

import (
	"context"
	"strings"
	"sync"

	"github.com/apache/kvrocks-controller/consts"
)

type Entry struct {
	Key   string
	Value []byte
}

type Engine interface {
	ID() string
	Leader() string
	LeaderChange() <-chan bool
	IsReady(ctx context.Context) bool

	Get(ctx context.Context, key string) ([]byte, error)
	Exists(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key string, value []byte) error
	Delete(ctx context.Context, key string) error
	List(ctx context.Context, prefix string) ([]Entry, error)

	Close() error
}

type Mock struct {
	mu     sync.Mutex
	values map[string]string
}

func NewMock() *Mock {
	return &Mock{
		values: make(map[string]string),
	}
}

func (m *Mock) Get(ctx context.Context, key string) ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.values[key]
	if !ok {
		return nil, consts.ErrNotFound
	}
	return []byte(v), nil
}

func (m *Mock) Exists(ctx context.Context, key string) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.values[key]
	return ok, nil
}

func (m *Mock) Set(ctx context.Context, key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.values[key] = string(value)
	return nil
}

func (m *Mock) Delete(ctx context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.values, key)
	return nil
}

func (m *Mock) List(ctx context.Context, prefix string) ([]Entry, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var entries []Entry
	for k, v := range m.values {
		if k == prefix || (len(k) > len(prefix) && k[len(prefix)] == '/') {
			entries = append(entries, Entry{
				Key:   strings.TrimPrefix(k, prefix+"/"),
				Value: []byte(v),
			})
		}
	}
	return entries, nil
}

func (m *Mock) Close() error {
	return nil
}

func (m *Mock) ID() string {
	return "mock_store_engine"
}

func (m *Mock) Leader() string {
	return "mock_store_engine"
}

func (m *Mock) LeaderChange() <-chan bool {
	return make(chan bool)
}

func (m *Mock) IsReady(ctx context.Context) bool {
	return true
}

var _ Engine = (*Mock)(nil)