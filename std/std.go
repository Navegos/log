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

package std

import (
	"bytes"
	stdlog "log"

	"github.com/Navegos/log"
)

// NewLogger creates a standard library logger that writes to logger at the designated
// level. This is useful for providing loggers to libraries that only accept the standard
// library logger.
func NewLogger(logger log.Logger, level log.Level) *stdlog.Logger {
	return stdlog.New(&logWriter{
		// stdlogger.Print -> stdlogger.Output -> Write -> logger
		logger: logger.AddCallerSkip(3),
		level:  level,
	}, "", 0)
}

// logWriter is an io.Writer that doesn't really implement io.Writer correctly, but
// implements it correctly enough to satisfy the needs of stdlog.Logger's usage of
// io.Writer. Notably, stdlog.Logger:
//
// - does not use the bytes written return value
// - guarantees that each call to Write is a separate message
type logWriter struct {
	logger log.Logger
	level  log.Level
}

func (w *logWriter) Write(p []byte) (int, error) {
	msg := string(bytes.TrimSuffix(p, []byte("\n")))
	switch w.level {
	case log.LevelDebug:
		w.logger.Debug(msg)
	case log.LevelInfo:
		w.logger.Info(msg)
	case log.LevelWarn:
		w.logger.Warn(msg)
	case log.LevelError:
		w.logger.Error(msg)
	}
	return len(msg), nil
}
