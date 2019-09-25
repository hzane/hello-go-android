# hello-go-android

```bash
#!/usr/bin/env bash

export GO_ANDROID_TOOLCHAIN=/opt/arm64-clang
export PATH=$PATH:${GO_ANDROID_TOOLCHAIN}/bin

# Tell configure what tools to use.
# go android only supports gcc
TARGET_HOST=aarch64-linux-android
export AR=${TARGET_HOST}-ar
export AS=${TARGET_HOST}-gcc
export CC=${TARGET_HOST}-gcc
export CXX=${TARGET_HOST}-g++
export LD=${TARGET_HOST}-ld
export STRIP=${TARGET_HOST}-strip

# Tell configure what flags Android requires.
export CFLAGS="-fPIE -fPIC"
export LDFLAGS="-pie"

export CGO_ENABLED=1
export GOARCH=arm64
export GOOS=android

go build -buildmode=pie
```
