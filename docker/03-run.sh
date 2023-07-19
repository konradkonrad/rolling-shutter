#!/usr/bin/env bash
set -xe

if docker compose ls >/dev/null 2>&1; then
  DC="docker compose"
else
  DC=docker-compose
fi

DC_ENV="${DC_ENV:-dev}"
if [[ "${DC_ENV}" = "single" ]]; then
	DC="${DC} -f docker-compose.single.yml"
else
	DC="${DC} -f docker-compose.yml"
fi

echo "Starting entire system"
$DC --profile dev up -d
sleep 5
$DC ps
