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

package logr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Navegos/log/logtest"
)

func TestGetLogger(t *testing.T) {
	logr := New(logtest.Scoped(t))

	t.Run("from the root", func(t *testing.T) {
		logger, ok := GetLogger(logr)
		assert.True(t, ok)
		assert.NotNil(t, logger)
	})

	t.Run("from a named sub-logger", func(t *testing.T) {
		logger, ok := GetLogger(logr.WithName("foobar"))
		assert.True(t, ok)
		assert.NotNil(t, logger)
	})
}
