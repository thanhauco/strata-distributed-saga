# The Transactional Outbox Pattern

Directly updating a database and publishing an event to a message broker introduces the **Dual-Write Problem**:
- If DB write succeeds but network fails before publish, the message is lost.
- If publish succeeds but DB transaction aborts, phantom messages pollute downstream consumers.

**Strata Solution**:
We write the business state and the event record atomically into Amazon DynamoDB inside a single `TransactWriteItems` request. DynamoDB Streams then guarantees at-least-once Change Data Capture (CDC) delivery to Amazon EventBridge.
