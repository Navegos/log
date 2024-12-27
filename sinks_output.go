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

package log

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Navegos/log/internal/sinkcores/outputcore"
	"github.com/Navegos/log/internal/stderr"
	"github.com/Navegos/log/output"
)

type outputSink struct {
	development bool
}

func (s *outputSink) Name() string { return "OutputSink" }

func (s *outputSink) build() (zapcore.Core, error) {
	w, err := stderr.Open()
	if err != nil {
		return nil, err
	}

	level := zap.NewAtomicLevelAt(Level(os.Getenv(EnvLogLevel)).Parse())
	format := output.ParseFormat(os.Getenv(EnvLogFormat))

	if s.development {
		format = output.FormatConsole
	}

	sampling, err := parseSamplingConfig()
	if err != nil {
		return nil, err
	}

	overrides, err := parseOverrides()
	if err != nil {
		return nil, err
	}

	return outputcore.NewCore(w, level, format, sampling, overrides, s.development), nil
}

// update is a no-op because outputSink cannot be changed live.
func (s *outputSink) update(updated SinksConfig) error { return nil }

func parseSamplingConfig() (config zap.SamplingConfig, err error) {
	if val, set := os.LookupEnv(EnvLogSamplingInitial); set {
		config.Initial, err = strconv.Atoi(val)
		if err != nil {
			err = fmt.Errorf("SRC_LOG_SAMPLING_INITIAL is invalid: %w", err)
			return
		}
	} else {
		config.Initial = 100
	}

	if val, set := os.LookupEnv(EnvLogSamplingInitial); set {
		config.Thereafter, err = strconv.Atoi(val)
		if err != nil {
			err = fmt.Errorf("SRC_LOG_SAMPLING_THEREAFTER is invalid: %w", err)
			return
		}
	} else {
		config.Thereafter = 100
	}

	return
}

func parseOverrides() ([]outputcore.Override, error) {
	raw := os.Getenv(EnvLogScopeLevel)
	var overrides []outputcore.Override
	for _, kv := range strings.Split(raw, ",") {
		if kv == "" {
			continue
		}

		p := strings.SplitN(kv, "=", 2)
		if len(p) != 2 {
			return nil, fmt.Errorf("%s=%q is invalid", EnvLogScopeLevel, raw)
		}
		overrides = append(overrides, outputcore.Override{
			Scope: p[0],
			Level: Level(p[1]).Parse(),
		})
	}
	return overrides, nil
}
