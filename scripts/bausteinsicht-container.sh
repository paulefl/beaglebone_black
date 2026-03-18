#!/bin/bash
# bausteinsicht-container.sh — run Bausteinsicht via Podman container
# Clones https://github.com/docToolchain/Bausteinsicht.git, builds the image,
# and runs it with the project model file mounted.
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_URL="https://github.com/docToolchain/Bausteinsicht.git"
REPO_DIR="$SCRIPT_DIR/.bausteinsicht-repo"
IMAGE_NAME="bausteinsicht:local"
MODEL="$SCRIPT_DIR/../arch/model/beaglebone_black.jsonc"
MODEL_DIR="$(cd "$(dirname "$MODEL")" && pwd)"
MODEL_IN_CONTAINER="/model/$(basename "$MODEL")"

CMD="${1:-}"

if [[ -z "$CMD" ]]; then
  echo "Usage: $0 <command> [args...]"
  echo "Example: $0 validate"
  echo "         $0 export --output /output"
  exit 1
fi
shift

# 1. Clone or update repository
if [[ ! -d "$REPO_DIR/.git" ]]; then
  echo "Cloning Bausteinsicht repo..."
  git clone "$REPO_URL" "$REPO_DIR"
else
  echo "Updating Bausteinsicht repo..."
  git -C "$REPO_DIR" pull --ff-only 2>/dev/null || true
fi

# 2. Build container image from Dockerfile in repo
echo "Building Bausteinsicht container image..."
podman build -t "$IMAGE_NAME" "$REPO_DIR"

# 3. Run container with model directory mounted
echo "Running: bausteinsicht $CMD --model $MODEL_IN_CONTAINER $*"
podman run --rm \
  -v "$MODEL_DIR:/model:ro" \
  "$IMAGE_NAME" \
  bausteinsicht "$CMD" --model "$MODEL_IN_CONTAINER" "$@"
