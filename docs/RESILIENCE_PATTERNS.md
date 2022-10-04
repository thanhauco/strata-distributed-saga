# Distributed Resilience Patterns

Strata incorporates three core stability design patterns:
1. **Exponential Backoff with Full Jitter**: Neutralizes the Thundering Herd problem against downstream APIs.
2. **Circuit Breaker**: Prevents cascading failures when payment gateways suffer partial brownouts.
3. **Bulkhead Isolation**: Partitions resource pools so that slow logistics carrier calls cannot starve memory.
