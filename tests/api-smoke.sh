#!/usr/bin/env bash
set -euo pipefail
BASE_URL="${BASE_URL:-http://localhost:8080}"
curl --fail --silent "$BASE_URL/health"; echo
curl --fail --silent "$BASE_URL/api/v1/servers"; echo
R=$(curl --fail --silent -X POST "$BASE_URL/api/v1/servers" -H 'Content-Type: application/json' -d '{"name":"smoke-test-server","hostname":"smoke.example.internal","ip_address":"10.20.30.40","environment":"test","provider":"aws"}')
echo "$R"
ID=$(printf '%s' "$R" | sed -n 's/.*"id":\([0-9]*\).*/\1/p')
test -n "$ID"
curl --fail --silent "$BASE_URL/api/v1/servers/$ID"; echo
CODE=$(curl --silent -o /dev/null -w '%{http_code}' -X DELETE "$BASE_URL/api/v1/servers/$ID")
test "$CODE" = 204
echo "Smoke test passed."
