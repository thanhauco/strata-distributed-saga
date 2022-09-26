package telemetry

import "fmt"

type TraceContext struct {
	TraceID string
	SpanID  string
}

func FormatTraceParent(traceID, spanID string) string {
	return fmt.Sprintf("00-%s-%s-01", traceID, spanID)
}

func AddTraceAnnotation(key, value string) {
	// Hook for AWS X-Ray subsegment metadata
}
