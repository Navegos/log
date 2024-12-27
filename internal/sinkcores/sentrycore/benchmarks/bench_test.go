// Copyright (C) 2025 @Navegos & @DevelVitorF Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package actor provides the structures for representing an actor who has
// access to resources.

package benchmarks_test

import (
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/Navegos/log"
	"github.com/Navegos/log/internal/configurable"
	"github.com/Navegos/log/internal/sinkcores/sentrycore"
	"github.com/Navegos/log/logtest"
)

// BenchmarkWithSentry-10           2253642              5205 ns/op            9841 B/op         87 allocs/op
func BenchmarkWithSentry(b *testing.B) {
	logger, _, _ := newTestLogger(b)

	err := errors.New("foobar")
	for n := 0; n < b.N; n++ {
		logger.With(log.Error(err)).Warn("msg", log.Int("key", 5))
	}
}

// BenchmarkWithoutSentry-10        2656189              4537 ns/op            6334 B/op         44 allocs/op
func BenchmarkWithoutSentry(b *testing.B) {
	logger, _ := logtest.Captured(b)
	err := errors.New("foobar")
	for n := 0; n < b.N; n++ {
		logger.With(log.Error(err), log.Int("key", 5)).Warn("msg")
	}
}

func newTestLogger(t testing.TB) (log.Logger, *sentrycore.TransportMock, func()) {
	transport := &sentrycore.TransportMock{}
	client, err := sentry.NewClient(sentry.ClientOptions{Transport: transport})
	require.NoError(t, err)

	core := sentrycore.NewCore(sentry.NewHub(client, sentry.NewScope()))

	cl := configurable.Cast(logtest.Scoped(t))

	return cl.WithCore(func(c zapcore.Core) zapcore.Core {
			return zapcore.NewTee(c, core)
		}),
		transport,
		func() {
			err := core.Sync()
			require.NoError(t, err)
		}
}
