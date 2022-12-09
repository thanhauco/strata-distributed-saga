# Chaos Engineering and Fault Injection

To validate distributed resilience, Strata subjects its handlers to simulated failure modes:
1. **Network Blackhole (504)**: Verifies that Step Functions retry policies do not trigger duplicate orders.
2. **Crash-After-Debit**: Verifies that the Transactional Outbox ensures message dispatch even if Lambda crashes mid-execution.
3. **Partition Isolation**: Verifies that the circuit breaker trips within 3 consecutive errors.
