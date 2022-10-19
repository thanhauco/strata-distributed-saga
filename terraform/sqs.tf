resource "aws_sqs_queue" "saga_dlq" {
  name                      = "strata-saga-dlq.fifo"
  fifo_queue                = true
  content_based_deduplication = true
  message_retention_seconds = 1209600 # 14 days
  kms_master_key_id         = "alias/aws/sqs"
}
