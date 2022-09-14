package main

import (
	"sync/atomic"
	"time"
)

// SpanConfig is a group of options for a Span.
type SpanConfig struct {
	timestamp  time.Time
	newRoot    bool
	stackTrace bool
}

// SpanEndOption applies an option to a SpanConfig. These options are
// applicable only when the span is ended.
type SpanEndOption interface {
	applySpanEnd(SpanConfig) SpanConfig
}

// Span is the individual component of a trace. It represents a single named
// and timed operation of a workflow that is traced. A Tracer is used to
// create a Span and it is then up to the operation the Span represents to
// properly end the Span when the operation itself ends.
//
// Warning: methods may be added to this interface in minor releases.
type Span interface {
	// End completes the Span. The Span is considered complete and ready to be
	// delivered through the rest of the telemetry pipeline after this method
	// is called. Therefore, updates to the Span are not allowed after this
	// method has been called.
	End(options ...SpanEndOption)

	// IsRecording returns the recording state of the Span. It will return
	// true if the Span is active and events can be recorded.
	IsRecording() bool

	// SetName sets the Span name.
	SetName(name string)
}

// tracer is a placeholder for a trace.Tracer.
//
// All Tracer functionality is forwarded to a delegate once configured.
// Otherwise, all functionality is forwarded to a NoopTracer.
type tracer struct {
	name     string
	delegate atomic.Value
}

// nonRecordingSpan is a minimal implementation of a Span that wraps a
// SpanContext. It performs no operations other than to return the wrapped
// SpanContext.
type nonRecordingSpan struct {
	tracer *tracer
}

var _ Span = nonRecordingSpan{}

// IsRecording always returns false.
func (nonRecordingSpan) IsRecording() bool { return false }

// SetError does nothing.
func (nonRecordingSpan) SetError(bool) {}

// End does nothing.
func (nonRecordingSpan) End(...SpanEndOption) {}

// SetName does nothing.
func (nonRecordingSpan) SetName(string) {}

func main() {

}
