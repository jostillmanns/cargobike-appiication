#!/bin/bash

set -ex

export TMP_DIR=$(ssh root@veedelvelo.de mktemp -d)
export CGO_ENABLED=0

ssh root@veedelvelo.de systemctl stop traefik || true
ssh root@veedelvelo.de systemctl stop cargobike || tre

go build
scp cargobike root@veedelvelo.de:/usr/local/bin/cargobike

ssh root@veedelvelo.de wget -q -O $TMP_DIR/traefik.tar.gz https://github.com/traefik/traefik/releases/download/v2.4.8/traefik_v2.4.8_linux_amd64.tar.gz
ssh root@veedelvelo.de tar Cxvf $TMP_DIR $TMP_DIR/traefik.tar.gz
ssh root@veedelvelo.de cp $TMP_DIR/traefik /usr/local/bin/traefik

ssh root@veedelvelo.de mkdir -p /etc/traefik
scp traefik.yaml root@veedelvelo.de:/etc/traefik.yaml
scp provider.yaml root@veedelvelo.de:/etc/traefik/provider.yaml

scp cargobike.service root@veedelvelo.de:/etc/systemd/system/cargobike.service
scp traefik.service root@veedelvelo.de:/etc/systemd/system/traefik.service
ssh root@veedelvelo.de systemctl daemon-reload
ssh root@veedelvelo.de systemctl enable cargobike
ssh root@veedelvelo.de systemctl enable traefik

ssh root@veedelvelo.de systemctl restart traefik
ssh root@veedelvelo.de systemctl restart cargobike

rm -rf $TMP_DIR
