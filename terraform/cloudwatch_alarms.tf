resource "aws_cloudwatch_metric_alarm" "saga_failures" {
  alarm_name          = "strata-high-saga-failure-rate"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = 2
  metric_name         = "ExecutionsFailed"
  namespace           = "AWS/States"
  period              = 60
  statistic           = "Sum"
  threshold           = 5
  alarm_description   = "Triggered when Step Functions Saga executions fail repeatedly"
}
