#!/usr/bin/env bash
set -xe

BB="docker run --rm -v$(pwd)/data:/data -w / busybox"
if docker compose ls >/dev/null 2>&1; then
  DC="docker compose"
else
  DC=docker-compose
fi

DC_ENV="${DC_ENV:-dev}"
if [[ "${DC_ENV}" = "single" ]]; then
	DC="${DC} -f docker-compose.single.yml"
	cmds="snapshot ${HOSTNAME}"
else
	DC="${DC} -f docker-compose.yml"
	cmds="snapshot keyper-0 keyper-1 keyper-2"
fi

$DC stop db
$DC rm -f db

${BB} rm -rf data/db

$DC up -d db
sleep 40

for cmd in ${cmds}; do
  $DC exec db createdb -U postgres $cmd
  $DC run --rm --no-deps $cmd initdb --config /config/${cmd}.toml
done

$DC stop db
