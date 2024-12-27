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

package log

import (
	"testing"

	"github.com/getsentry/sentry-go"
	"github.com/stretchr/testify/assert"
)

func TestNewSentrySink(t *testing.T) {
	t.Run("defaults are applied", func(t *testing.T) {
		s := NewSentrySink()
		ss, ok := s.(*sentrySink)
		assert.True(t, ok)
		assert.Equal(t, 0.1, ss.SentrySink.SampleRate)
	})

	t.Run("values are set", func(t *testing.T) {
		s := NewSentrySinkWith(SentrySink{ClientOptions: sentry.ClientOptions{
			Dsn:   "http://test:test@do.no.exist/123",
			Debug: true,
		}})
		ss, ok := s.(*sentrySink)
		assert.True(t, ok)
		assert.Equal(t, "http://test:test@do.no.exist/123", ss.SentrySink.Dsn)
		assert.Equal(t, true, ss.SentrySink.Debug)
	})
}

func TestSentrySinkUpdate(t *testing.T) {
	t.Run("sink is updated with values", func(t *testing.T) {
		s := NewSentrySinkWith(SentrySink{ClientOptions: sentry.ClientOptions{
			Dsn: "http://test:test@do.no.exist/123",
		}})
		_, err := s.build()
		assert.Nil(t, err)

		err = s.update(SinksConfig{
			&SentrySink{
				ClientOptions: sentry.ClientOptions{
					Dsn:        "",
					SampleRate: 0.3333,
				},
			}})

		assert.Nil(t, err)

		ss, ok := s.(*sentrySink)
		assert.True(t, ok)

		assert.Equal(t, "", ss.Dsn)
		assert.Equal(t, 0.3333, ss.SampleRate)
	})
}
