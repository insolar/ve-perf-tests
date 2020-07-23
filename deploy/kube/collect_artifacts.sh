#!/usr/bin/env bash
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
KUBECTL=${KUBECTL:-"kubectl"}
ARTIFACTS_DIR=${ARTIFACTS_DIR:-"/tmp/insolar"}
LOG_DIR="$ARTIFACTS_DIR/logs"
NODES_COUNT=5
KUB_NAMESPACE=""

save_logs_to_files() {
  LOG_DIR="$ARTIFACTS_DIR/logs"
  rm -rf "$LOG_DIR"
  mkdir -p "$LOG_DIR"

  for ((i=0; i < ${NODES_COUNT}; i++))
  do
     $KUBECTL -n ${KUB_NAMESPACE} logs virtual-$i >"$LOG_DIR/virtual-$i"
  done

  $KUBECTL -n ${KUB_NAMESPACE} logs bootstrap >"$LOG_DIR/bootstrap"
}

save_logs_to_files
echo "Logs saved to $LOG_DIR"
