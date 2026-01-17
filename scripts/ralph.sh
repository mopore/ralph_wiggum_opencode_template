#!/usr/bin/env bash
set -euo pipefail

ITERATIONS="${1:-}"
if [[ -z "$ITERATIONS" ]]; then
  echo "Usage: $0 <iterations>"
  exit 1
fi

for ((i=1; i<=ITERATIONS; i++)); do
  echo "Iteration $i"
  echo "----------------------------------------"

  cmd=(opencode run --agent build --file plans/prd.json --file progress.txt)

  if [[ -n "${OPENCODE_ATTACH:-}" ]]; then
    cmd+=(--attach "$OPENCODE_ATTACH")
  fi

  result=$("${cmd[@]}" < plans/prompt.md)

  echo "$result"

  if [[ "$result" == *"<promise>COMPLETE</promise>"* ]]; then
    echo "PRD complete, exiting."
    exit 0
  fi

  echo
  echo
  sleep 1
done
