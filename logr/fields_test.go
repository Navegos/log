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

	"github.com/Navegos/log"
)

func TestToLogFields(t *testing.T) {
	fields := toLogFields([]any{
		"hello", "world",
		"goodbye", "bob",
		"lucky_number", 3,
	})
	assert.Equal(t, fields, []log.Field{
		log.String("hello", "world"),
		log.String("goodbye", "bob"),
		log.Int("lucky_number", 3),
	})
}
