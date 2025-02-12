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

package main

import (
	"os"
	"time"

	"github.com/Navegos/log"
)

func main() {
	liblog := log.Init(log.Resource{
		Name: "logexample",
	})
	defer liblog.Sync()

	l := log.Scoped("foo")

	// print diagnostics
	config := []log.Field{}
	for _, k := range []string{
		log.EnvDevelopment,
		log.EnvLogFormat,
		log.EnvLogLevel,
		log.EnvLogScopeLevel,
		log.EnvLogSamplingInitial,
		log.EnvLogSamplingThereafter,
	} {
		config = append(config, log.String(k, os.Getenv(k)))
	}
	l.Info("configuration", config...)

	// sample message
	l.Warn("hello world!", log.Time("now", time.Now()))
}
