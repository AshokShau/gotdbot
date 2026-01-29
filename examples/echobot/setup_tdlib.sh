#!/usr/bin/env bash
set -euo pipefail

TD_DIR="td"
OUT_DIR="$(pwd)/"

rm -rf "$TD_DIR" "$OUT_DIR"

git clone --depth=1 https://github.com/tdlib/td.git "$TD_DIR"

mkdir -p "$TD_DIR/build" "$OUT_DIR"
cd "$TD_DIR/build"

cmake .. \
  -DCMAKE_BUILD_TYPE=Release \
  -DBUILD_SHARED_LIBS=ON \
  -DCMAKE_LIBRARY_OUTPUT_DIRECTORY="$OUT_DIR"

cmake --build . -- -j"$(nproc)"

echo "libtdjson.so built at: $OUT_DIR"
