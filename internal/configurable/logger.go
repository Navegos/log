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

package configurable

import (
	"go.uber.org/zap/zapcore"

	"github.com/Navegos/log"
)

// Logger exposes internal APIs that must be implemented on
// github.com/Navegos/log.zapAdapter
type Logger interface {
	log.Logger

	// WithCore is an internal API used to allow packages like logtest to hook into
	// underlying zap logger's core.
	WithCore(func(c zapcore.Core) zapcore.Core) log.Logger
}

// Cast provides a configurable logger API for testing purposes.
func Cast(l log.Logger) Logger { return l.(Logger) }
