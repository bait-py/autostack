#!/bin/bash

# Test script for AutoStack with automated inputs
# This simulates user inputs for testing

echo "=== Testing LAMP Stack Creation ==="
echo ""

# Create input file with responses
# Format: env vars (4), ports (3), confirm, auto-start
cat << EOF | ./autostack create lamp
custom_root_pass
testdb
dbuser
dbpass123
8888
3307
8889
y
n
EOF

echo ""
echo "=== Generated files in lamp-stack/ ==="
ls -la lamp-stack/

echo ""
echo "=== Checking docker-compose.yml for custom ports ==="
grep -E "PORT|8888|3307|8889" lamp-stack/docker-compose.yml || echo "Ports not found in docker-compose.yml"

echo ""
echo "=== Checking README.md for custom ports ==="
grep -E "8888|3307|8889" lamp-stack/README.md || echo "Ports not found in README.md"

echo ""
echo "=== Test completed ==="
