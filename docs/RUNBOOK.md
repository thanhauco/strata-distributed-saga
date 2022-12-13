# Production Operations Runbook

## 1. High Failure Rate Alert (strata-high-saga-failure-rate)
1. Query active Step Functions executions:
   ```bash
   aws stepfunctions list-executions --state-machine-arn <ARN> --status-filter FAILED
   ```
2. Inspect the failure cause:
   ```bash
   saga-cli -saga=<saga_id> -action=status
   ```
3. If poison pills are detected in the SQS FIFO DLQ, run the redrive script:
   ```bash
   ./scripts/replay_dlq.sh
   ```
