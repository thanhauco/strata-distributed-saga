# Distributed Idempotency Specification

Strata achieves exactly-once execution semantics by enforcing:
1. **Deterministic Hashing**: SHA-256(SagaID + StepName + PayloadHash).
2. **Conditional Locking**: DynamoDB `attribute_not_exists(idempotency_key)`.
3. **Response Caching**: Successful responses stored alongside the key to instantly fulfill retransmitted requests.
