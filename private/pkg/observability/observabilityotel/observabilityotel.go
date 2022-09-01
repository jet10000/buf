// Copyright 2020-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package observabilityotel

import (
	"context"

	"github.com/bufbuild/buf/private/pkg/observability"
	"go.opentelemetry.io/otel/trace"
)

type TracerProvider interface {
	Tracer(string, ...trace.TracerOption) trace.Tracer
	Shutdown(context.Context) error
}

func NewTracerProviderCloser(tracerProvider TracerProvider) observability.TracerProviderCloser {
	return newTracerProviderCloser(tracerProvider)
}
