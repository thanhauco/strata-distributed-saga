resource "aws_sfn_state_machine" "order_saga" {
  name     = "strata-order-saga"
  role_arn = aws_iam_role.lambda_exec.arn

  definition = templatefile("${path.module}/../statemachine/order_saga.asl.json", {
    payment_reserve_lambda_arn   = "arn:aws:lambda:us-east-1:123456789012:function:payment-reserve"
    inventory_reserve_lambda_arn = "arn:aws:lambda:us-east-1:123456789012:function:inventory-reserve"
    shipping_dispatch_lambda_arn = "arn:aws:lambda:us-east-1:123456789012:function:shipping-dispatch"
    payment_refund_lambda_arn    = "arn:aws:lambda:us-east-1:123456789012:function:payment-refund"
    inventory_release_lambda_arn = "arn:aws:lambda:us-east-1:123456789012:function:inventory-release"
  })
}
