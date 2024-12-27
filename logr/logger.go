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
	"github.com/go-logr/logr"

	"github.com/Navegos/log"
)

// GetLogger retrieves the underlying log.Logger. If no Logger is found,
// a Logger scoped to 'logr' is returned. The second return value can be
// checked if such a Logger was created.
func GetLogger(l logr.Logger) (log.Logger, bool) {
	sink, ok := l.GetSink().(*LogSink)
	if !ok {
		return log.Scoped("logr"), false
	}
	return sink.Logger, true
}
