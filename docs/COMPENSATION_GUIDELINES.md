# Microservice Compensation Guidelines

Compensating transactions in a Saga must strictly adhere to three properties:
1. **Idempotence**: Executing a refund twice must have the same side effect as executing it once.
2. **Commutativity**: Reversing step B before or after reversing step A must yield consistent state.
3. **Eventual Inevitability**: Compensations must never fail permanently. Transient failures must retry with exponential backoff until success.
