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

package log_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Navegos/log"
	"github.com/Navegos/log/internal/globallogger"
	"github.com/Navegos/log/internal/otelfields"
	"github.com/Navegos/log/logtest"
)

func newTestLogger(t *testing.T) (log.Logger, func() logtest.CapturedLogs) {
	logger, exportLogs := logtest.Captured(t)
	assert.NotNil(t, logger)

	// HACK: If in devmode, the attributes namespace does not get added, but we want to
	// test that behaviour here so we add it back.
	if globallogger.DevMode() {
		logger = logger.With(otelfields.AttributesNamespace)
	}

	return logger, exportLogs
}

func TestLogger(t *testing.T) {

	logger, exportLogs := newTestLogger(t)

	logger.Debug("a debug message") // 0

	logger = logger.With(log.String("some", "field"))

	logger.Info("hello world", log.String("hello", "world")) // 1

	logger = logger.WithTrace(log.TraceContext{TraceID: "1234abcde"})
	logger.Info("goodbye", log.String("world", "hello")) // 2
	logger.Warn("another message")                       // 3

	logger.Error("object of fields", // 4
		log.Object("object",
			log.String("field1", "value"),
			log.String("field2", "value"),
		))

	logs := exportLogs()
	assert.Len(t, logs, 5)
	for _, l := range logs {
		assert.Equal(t, "TestLogger", l.Scope) // scope is always applied
	}

	// Nested fields should be in attributes
	assert.Equal(t, map[string]interface{}{
		"some":  "field",
		"hello": "world",
	}, logs[1].Fields["Attributes"])

	// TraceId should be in root, everything else in attributes
	assert.Equal(t, "1234abcde", logs[2].Fields["TraceId"])
	assert.Equal(t, map[string]interface{}{
		"some":  "field",
		"world": "hello",
	}, logs[2].Fields["Attributes"])

	// Nested fields should be in attributes
	assert.Equal(t, map[string]interface{}{
		"some": "field",
		"object": map[string]interface{}{
			"field1": "value",
			"field2": "value",
		},
	}, logs[4].Fields["Attributes"])
}

func TestWithTrace(t *testing.T) {
	logger, exportLogs := newTestLogger(t)

	logger = logger.WithTrace(log.TraceContext{})
	logger.Info("should not have trace") // 0

	logger = logger.WithTrace(log.TraceContext{TraceID: "1"})
	logger.Info("should have trace 1") // 1

	logger = logger.WithTrace(log.TraceContext{TraceID: "2"})
	logger.Info("should have trace 2") // 2

	logs := exportLogs()
	assert.NotContains(t, logs[0].Fields, "TraceId")
	assert.Equal(t, "1", logs[1].Fields["TraceId"])
	assert.Equal(t, "2", logs[2].Fields["TraceId"])
}