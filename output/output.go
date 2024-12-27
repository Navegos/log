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

package output

// Format configures Navegos/log output encoding.
type Format string

const (
	// FormatJSON encodes log entries to a machine-readable, OpenTelemetry-structured
	// format.
	FormatJSON Format = "json"
	// FormatJSONGCP encodes log entries to a machine-readable, GCP-structured format.
	// It's similar to OpenTelemetry-structured format, but the severity field
	// complies with https://cloud.google.com/logging/docs/structured-logging#special-payload-fields
	FormatJSONGCP Format = "json_gcp"
	// FormatConsole encodes log entries to a human-readable format.
	FormatConsole Format = "console"
)

// ParseFormat parses the given format string as a supported output format, while
// trying to maintain some degree of back-compat with the intent of previously supported
// log formats.
func ParseFormat(format string) Format {
	switch format {
	case string(FormatJSONGCP):
		return FormatJSONGCP

	// True 'logfmt' has significant limitations around certain field types:
	// https://github.com/jsternberg/zap-logfmt#limitations so since it implies a
	// desire for a somewhat structured format, we interpret it as OutputJSON.
	case string(FormatJSON), "logfmt":
		return FormatJSON

	// The previous 'condensed' format is optimized for local dev, so it serves the
	// same purpose as OutputConsole
	case string(FormatConsole), "condensed":
		return FormatConsole
	}

	// Fall back to JSON output
	return FormatJSON
}
