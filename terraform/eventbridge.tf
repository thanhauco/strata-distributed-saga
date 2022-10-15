resource "aws_cloudwatch_event_bus" "saga_bus" {
  name = "strata-saga-event-bus"
}

resource "aws_cloudwatch_event_rule" "saga_events_rule" {
  name           = "strata-capture-saga-events"
  event_bus_name = aws_cloudwatch_event_bus.saga_bus.name

  event_pattern = jsonencode({
    source = ["strata.order", "strata.payment", "strata.inventory"]
  })
}
