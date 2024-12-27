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

import "go.uber.org/zap/zapcore"

// Zap levels are int8 - make sure we stay in bounds.  logr itself should
// ensure we never get negative values.
//
// Source: https://github.com/go-logr/zapr/blob/48df242fffb25049c72e208aea4826177ff5fe8e/zapr.go#L196
func toZapLevel(lvl int) zapcore.Level {
	if lvl > 127 {
		lvl = 127
	}
	// zap levels are inverted.
	return 0 - zapcore.Level(lvl)
}
