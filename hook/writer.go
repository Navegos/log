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

package hook

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Navegos/log"
	"github.com/Navegos/log/internal/configurable"
	"github.com/Navegos/log/internal/sinkcores/outputcore"
	"github.com/Navegos/log/output"
)

type writerSyncerAdapter struct{ io.Writer }

func (writerSyncerAdapter) Sync() error { return nil }

// Writer hooks receiver to rendered log output at level in the requested format,
// typically one of 'json' or 'console'.
func Writer(logger log.Logger, receiver io.Writer, level log.Level, format output.Format) log.Logger {
	cl := configurable.Cast(logger)

	// Adapt to WriteSyncer in case receiver doesn't implement it
	var writeSyncer zapcore.WriteSyncer
	if ws, ok := receiver.(zapcore.WriteSyncer); ok {
		writeSyncer = ws
	} else {
		writeSyncer = writerSyncerAdapter{receiver}
	}

	core := outputcore.NewCore(writeSyncer, level.Parse(), format, zap.SamplingConfig{}, nil, false)
	return cl.WithCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(c, core)
	})
}
