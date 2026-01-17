#!/usr/bin/env bash
set -euo pipefail

result=$(opencode run --agent build --file plans/prd.json --file progress.txt < plans/prompt.md)
echo "$result"
