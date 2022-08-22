# Saga State Machine Specification

```mermaid
stateDiagram-v2
    [*] --> ReservePayment
    ReservePayment --> ReserveInventory : Success
    ReservePayment --> SagaFailed : Payment Failed

    ReserveInventory --> DispatchShipping : Success
    ReserveInventory --> CompensatePayment : Inventory Depleted

    DispatchShipping --> SagaCompleted : Success
    DispatchShipping --> CompensateAll : Dispatch Failed

    CompensateAll --> SagaFailed
    CompensatePayment --> SagaFailed
```
