#!/usr/bin/env bash
set -euo pipefail

ITERATIONS="${1:-}"
if [[ -z "$ITERATIONS" ]]; then
  echo "Usage: $0 <iterations>"
  exit 1
fi

# Tip: for faster repeated runs you can start a headless server in another terminal:
#   opencode serve
# and then set OPENCODE_ATTACH=http://localhost:4096

ATTACH_FLAG=()
if [[ -n "${OPENCODE_ATTACH:-}" ]]; then
  ATTACH_FLAG=(--attach "$OPENCODE_ATTACH")
fi

for ((i=1; i<=ITERATIONS; i++)); do
  echo "Iteration $i"
  echo "----------------------------------------"

  # We pass the prompt via stdin (avoids shell escaping issues) and attach files explicitly.
  # If OpenCode hangs in CI, it's usually waiting for an 'ask' permission.
  # Tune opencode.json permissions if needed.
  result=$(opencode run --agent build --file plans/prd.json --file progress.txt "${ATTACH_FLAG[@]}" < plans/prompt.md)

  echo "$result"

  if [[ "$result" == *"<promise>COMPLETE</promise>"* ]]; then
    echo "PRD complete, exiting."
    exit 0
  fi

  echo
  echo
  sleep 1
done
