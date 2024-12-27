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

package hook_test

import (
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/require"
	"go.bobheadxi.dev/streamline/jq"
	"go.bobheadxi.dev/streamline/pipe"

	"github.com/Navegos/log"
	"github.com/Navegos/log/hook"
	"github.com/Navegos/log/logtest"
	"github.com/Navegos/log/output"
)

func TestWriter(t *testing.T) {
	logger, exportLogs := logtest.Captured(t)

	writer, stream := pipe.NewStream()
	hookedLogger := hook.Writer(logger, writer, log.LevelWarn, output.FormatJSON)

	hookedLogger.Debug("debug message")
	hookedLogger.Warn("warn message")
	hookedLogger.Error("error message")

	logger.Error("parent message")

	// done with writing
	writer.CloseWithError(nil)

	// hooked logger output - only warn and above, and messages logged to parent are not
	// included. We only get the messages because there's no easy way to mock the clock.
	hookedOutput, err := stream.WithPipeline(jq.Pipeline(".Body")).Lines()
	require.NoError(t, err)
	autogold.Expect([]string{`"warn message"`, `"error message"`}).Equal(t, hookedOutput)

	// parent logger output - should receive everything
	parentOutput := exportLogs().Messages()
	autogold.Expect([]string{
		"debug message", "warn message", "error message",
		"parent message",
	}).Equal(t, parentOutput)
}
