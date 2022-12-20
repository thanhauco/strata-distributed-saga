# Changelog

All notable changes to Strata Distributed Saga are documented in this file.

## [1.0.0] - 2022-12-20

### Added
- Complete AWS Step Functions ASL Saga state machine with compensation catch blocks.
- Microservice handlers in Go 1.18 for Orders, Payments, Inventory, and Shipping.
- Transactional Outbox pattern with DynamoDB Streams and EventBridge publisher.
- Deterministic SHA-256 idempotency manager with DynamoDB conditional writes.
- Resilience patterns: Full Jitter backoff, sliding-window Circuit Breaker, and Bulkhead concurrency limits.
- Terraform IaC for DynamoDB, EventBridge, SQS FIFO DLQ, IAM, and CloudWatch alarms.
- `saga-cli` administration tool, reconciliation worker, and chaos engineering scenarios.
- Production operations runbook and architecture specifications.
