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

package logtest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Navegos/log"
)

func TestExport(t *testing.T) {
	logger, exportLogs := Captured(t)
	assert.NotNil(t, logger)

	logger.Info("hello world", log.String("key", "value"))
	logger.Error("goodbye world", log.String("key", "value"))

	logs := exportLogs()
	assert.Len(t, logs, 2)
	assert.Equal(t, "TestExport", logs[0].Scope)    // test name is the scope
	assert.Equal(t, "hello world", logs[0].Message) // retains the message

	// In dev mode, attributes are not added, but custom fields are retained
	assert.Equal(t, map[string]interface{}{"key": "value"}, logs[0].Fields)

	// We can filter for entries
	assert.Len(t, logs.Filter(func(l CapturedLog) bool {
		return l.Level == log.LevelError
	}), 1)

	// We can assert the existence of an entry
	assert.False(t, logs.Contains(func(l CapturedLog) bool {
		return l.Level == log.LevelWarn
	}))
}
