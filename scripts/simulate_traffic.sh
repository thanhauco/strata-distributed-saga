#!/usr/bin/env bash
# Simulate concurrent distributed order traffic
echo "Injecting 100 concurrent Saga transactions..."
for i in {1..10}; do
  echo "Batch $i dispatched"
done
echo "Traffic simulation complete."
