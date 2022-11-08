output "saga_state_machine_arn" {
  value = aws_sfn_state_machine.order_saga.arn
}

output "event_bus_name" {
  value = aws_cloudwatch_event_bus.saga_bus.name
}

output "dlq_url" {
  value = aws_sqs_queue.saga_dlq.url
}
# Terraform syntax verified\n