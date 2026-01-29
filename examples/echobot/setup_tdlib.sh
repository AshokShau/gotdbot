#!/usr/bin/env bash
set -euo pipefail

TD_DIR="td"
BUILD_DIR="${TD_DIR}/build"
PROJECT_DIR="$(pwd)"


if [[ ! -d "${TD_DIR}/.git" ]]; then
  echo "Cloning TDLib..."
  git clone --depth=1 https://github.com/tdlib/td.git "${TD_DIR}"
fi

if [[ ! -f "${BUILD_DIR}/libtdjson.so" ]]; then
  echo "Building TDLib..."
  mkdir -p "${BUILD_DIR}"
  cd "${BUILD_DIR}"

  cmake .. \
    -DCMAKE_BUILD_TYPE=Release \
    -DBUILD_SHARED_LIBS=ON \
    -DTD_ENABLE_TESTS=OFF \
    -DTD_ENABLE_BENCHMARKS=OFF

  cmake --build . -- -j"$(nproc)"
else
  echo "TDLib already built ✔"
fi

cd "${PROJECT_DIR}"

REAL_LIB="$(readlink -f "${BUILD_DIR}/libtdjson.so")"

if [[ ! -f "${REAL_LIB}" ]]; then
  echo "Failed to resolve libtdjson real file"
  exit 1
fi

echo "Copying TDLib into project directory..."
cp -av "${REAL_LIB}" .
ln -sf "$(basename "${REAL_LIB}")" libtdjson.so

echo
echo "TDLib ready"
ls -l libtdjson.so*
