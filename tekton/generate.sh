#!/bin/sh
set -eu

. ./versions.env

curl -fsSL \
  "https://infra.tekton.dev/tekton-releases/operator/previous/${TEKTON_OPERATOR_VERSION}/release.yaml"

printf '\n---\n'

cat <<'YAML'
apiVersion: operator.tekton.dev/v1alpha1
kind: TektonConfig
metadata:
  name: config
spec:
  profile: all
YAML
