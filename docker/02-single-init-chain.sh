#!/usr/bin/env bash
set -xe

BB="docker run --rm -v$(pwd)/config:/config -v$(pwd)/data:/data -w / busybox"
if docker compose ls >/dev/null 2>&1; then
  # compose v2
  DC="docker compose"
else
  DC=docker-compose
fi

DC_ENV="${DC_ENV:-dev}"
if [[ "${DC_ENV}" = "single" ]]; then
	DC="${DC} -f docker-compose.single.yml"
else
	DC="${DC} -f docker-compose.yml"
	exit "Don't run this for local testing"
fi

$DC stop chain-{validator,sentry}
$DC rm -f chain-{validator,sentry}

${BB} rm -rf data/chain-{validator,sentry}
${BB} mkdir -p data/chain-{validator,sentry}/config
${BB} chmod -R a+rwX data/chain-{validator,sentry}/config
${BB} rm -rf data/deployments

if [[ "${HOSTNAME}" = "devkeyper-01" ]]; then
	# has geth as dependency
	$DC up deploy-contracts
fi

TM_P2P_PORT=26656
TM_RPC_PORT=26657

validator_cmd=chain-validator
sentry_cmd=chain-sentry

for destination in /data/chain-validator/config/ /data/chain-sentry/config/; do
  ${BB} cp -v /config/genesis.json "${destination}"
done

$DC run --rm --no-deps ${sentry_cmd} init \
   --root /chain \
   --blocktime 1 \
   --listen-address tcp://0.0.0.0:${TM_RPC_PORT} \
   --role sentry

# TODO: check if validator can have listen-address tcp://127.0.0.1...
$DC run --rm --no-deps ${validator_cmd} init \
   --root /chain \
   --genesis-keyper 0x440Dc6F164e9241F04d282215ceF2780cd0B755e \
   --blocktime 1 \
   --listen-address tcp://127.0.0.1:${TM_RPC_PORT} \
   --role validator

sed -i "/ValidatorPublicKey/c\ValidatorPublicKey = \"$(cat data/${validator_cmd}/config/priv_validator_pubkey.hex)\"" config/${HOSTNAME}.toml


seed_node=$(cat config/seed_nodes.txt)

sentry_cmd=chain-sentry
validator_cmd=chain-validator

validator_id=$(cat data/${validator_cmd}/config/node_key.json.id)
validator_node=${validator_id}@${validator_cmd}:${TM_P2P_PORT}
sentry_node=$(cat data/${sentry_cmd}/config/node_key.json.id)@${sentry_cmd}:${TM_P2P_PORT}

# set seed node for sentry
${BB} sed -i "/^persistent-peers =/c\persistent-peers = \"${seed_node}\"" data/${sentry_cmd}/config/config.toml
# set validator node for sentry
${BB} sed -i "/^private-peer-ids =/c\private-peer-ids = \"${validator_id}\"" data/${sentry_cmd}/config/config.toml
${BB} sed -i "/^unconditional-peer-ids =/c\unconditional-peer-ids = \"${validator_id}\"" data/${sentry_cmd}/config/config.toml
${BB} sed -i "/^external-address =/c\external-address = \"${HOST}:${TM_P2P_PORT}\"" data/${sentry_cmd}/config/config.toml

# set sentry node for validator
${BB} sed -i "/^persistent-peers =/c\persistent-peers = \"${sentry_node}\"" data/${validator_cmd}/config/config.toml
${BB} sed -i "/^external-address =/c\external-address = \"${validator_cmd}:${TM_P2P_PORT}\"" data/${validator_cmd}/config/config.toml

$DC up -d chain-{sentry,validator} ${HOSTNAME} 

echo "We need to wait for the chain to reach height >= 1"
sleep 5
echo "This will take a while..."


if [[ "${HOSTNAME}" = "devkeyper-01" ]]; then
$DC run --rm --no-deps --entrypoint /rolling-shutter chain-validator bootstrap \
  --deployment-dir /deployments/dockerGeth \
  --ethereum-url http://10.135.179.106:8545 \
  --shuttermint-url http://chain-sentry:${TM_RPC_PORT} \
  --signing-key 479968ffa5ee4c84514a477a8f15f3db0413964fd4c20b08a55fed9fed790fad
fi
$DC stop -t 30 chain-{sentry,validator} ${HOSTNAME} 
