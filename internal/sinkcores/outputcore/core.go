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

package outputcore

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Navegos/log/internal/encoders"
	"github.com/Navegos/log/output"
)

func NewCore(
	output zapcore.WriteSyncer,
	level zapcore.LevelEnabler,
	format output.Format,
	sampling zap.SamplingConfig,
	overrides []Override,
	development bool,
) zapcore.Core {
	newCore := func(level zapcore.LevelEnabler) zapcore.Core {
		return zapcore.NewCore(
			encoders.BuildEncoder(format, development),
			output,
			level,
		)
	}

	core := newOverrideCore(level, overrides, newCore)

	if sampling.Initial > 0 {
		return zapcore.NewSamplerWithOptions(core, time.Second, sampling.Initial, sampling.Thereafter)
	}
	return core
}
