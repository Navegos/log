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

package configurable_test

import (
	"testing"

	"github.com/Navegos/log"
	"github.com/Navegos/log/internal/configurable"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestCast(t *testing.T) {
	assert.NotPanics(t, func() {
		log.Init(log.Resource{Name: t.Name()})

		// Cast works
		cl := configurable.Cast(log.Scoped("foo"))

		// Core wrapping works
		_ = cl.WithCore(func(c zapcore.Core) zapcore.Core {
			return zapcore.NewTee(c, zapcore.NewNopCore())
		})
	})
}
