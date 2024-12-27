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

package std_test

import (
	"testing"

	"github.com/Navegos/log"
	"github.com/Navegos/log/logtest"
	"github.com/Navegos/log/std"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	root, export := logtest.Captured(t)

	l := std.NewLogger(root, log.LevelInfo)
	l.Println("foobar")

	l.SetPrefix("prefix: ")
	l.Println("baz")

	logs := export()
	assert.Len(t, logs, 2)

	assert.Equal(t, logs[0].Level, log.LevelInfo)
	assert.Equal(t, logs[0].Message, "foobar")

	assert.Equal(t, logs[1].Level, log.LevelInfo)
	assert.Equal(t, logs[1].Message, "prefix: baz")
}
