package config

import "os"

type Config struct {
	Region          string
	SagaTable       string
	IdempotencyTable string
	OutboxTable     string
	EventBusName    string
	DLQQueueURL     string
}

func Load() *Config {
	return &Config{
		Region:          getEnv("AWS_REGION", "us-east-1"),
		SagaTable:       getEnv("DYNAMO_SAGA_TABLE", "strata-saga-instances"),
		IdempotencyTable: getEnv("DYNAMO_IDEMPOTENCY_TABLE", "strata-idempotency-keys"),
		OutboxTable:     getEnv("DYNAMO_OUTBOX_TABLE", "strata-orders-outbox"),
		EventBusName:    getEnv("EVENTBRIDGE_BUS_NAME", "strata-saga-event-bus"),
		DLQQueueURL:     getEnv("SQS_DLQ_URL", ""),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
