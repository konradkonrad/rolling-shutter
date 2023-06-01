#!/usr/bin/env bash
set -xe

if docker compose ls >/dev/null 2>&1; then
  DC="docker compose"
else
  DC=docker-compose
fi

echo "Starting entire system"
$DC up -d --profile dev
sleep 5
$DC ps
