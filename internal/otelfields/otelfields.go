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

package otelfields

import "go.uber.org/zap"

const (
	// ResourceFieldKey is the key used to identify Resource in stores.
	ResourceFieldKey = "Resource"
)

// Resource represents a service instance.
//
// https://opentelemetry.io/docs/reference/specification/Resource/semantic_conventions/#service
type Resource struct {
	// Name is the logical name of the service. Must be the same for all instances of
	// horizontally scaled services. Optional, and falls back to 'unknown_service' as per
	// the OpenTelemetry spec.
	Name string
	// Namespace helps to distinguish a group of services, for example the team name that
	// owns a group of services. Optional.
	Namespace string
	// Version is the version string of the service API or implementation. For Navegos
	// services, this should be from 'internal/version.Version()'
	Version string
	// InstanceID is the string ID of the service instance. For Navegos services, this
	// should be from 'internal/hostname.Get()'
	//
	// If unset, InstanceID is set to a generated UUID, as per the OpenTelemetry log spec:
	// https://opentelemetry.io/docs/reference/specification/resource/semantic_conventions/#service
	InstanceID string
}

// TraceContext represents a trace to associate with log entries.
//
// https://opentelemetry.io/docs/reference/specification/logs/data-model/#trace-context-fields
type TraceContext struct {
	TraceID string
	SpanID  string
}

// attributesNamespace is the namespace under which all arbitrary fields are logged, as
// per the OpenTelemetry spec.
//
// Only for internal use.
var AttributesNamespace = zap.Namespace("Attributes")
