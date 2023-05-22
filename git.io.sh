#!/bin/sh

RE_BIN_DIR=$HOME/.re/bin

# check if env RE_ROOT is set
if [ -n "$RE_ROOT" ]; then
  RE_BIN_DIR=$RE_ROOT/.re/bin
fi

mkdir -p $RE_BIN_DIR

RE_ARCH_BIN=''
RE_BIN='re'

THISOS=$(uname -s)
ARCH=$(uname -m)

case $THISOS in
   Linux*)
      case $ARCH in
        arm64)
          RE_ARCH_BIN="re-linux-arm64"
          ;;
        aarch64)
          RE_ARCH_BIN="re-linux-arm64"
          ;;
        armv6l)
          RE_ARCH_BIN="re-linux-arm"
          ;;
        armv7l)
          RE_ARCH_BIN="re-linux-arm"
          ;;
        *)
          RE_ARCH_BIN="re-linux-amd64"
          ;;
      esac
      ;;
   Darwin*)
      case $ARCH in
        arm64)
          RE_ARCH_BIN="re-darwin-arm64"
          ;;
        *)
          RE_ARCH_BIN="re-darwin-amd64"
          ;;
      esac
      ;;
   Windows|MINGW64_NT*)
      RE_ARCH_BIN="re-windows-amd64.exe"
      RE_BIN="re.exe"
      ;;
esac

if [ -z "$RE_ARCH_BIN" ]; then
   echo "This script is not supported on $THISOS and $ARCH"
   exit 1
fi

echo "Installing re from: https://github.com/kevincobain2000/re/releases/latest/download/$RE_ARCH_BIN ..."
echo ""

curl -kLs https://github.com/kevincobain2000/re/releases/latest/download/$RE_ARCH_BIN -o $RE_BIN_DIR/$RE_BIN

chmod +x $RE_BIN_DIR/$RE_BIN

echo "Installed successfully to: $RE_BIN_DIR/$RE_BIN"
echo 'Add following to your shell config (.bashrc/.zshrc):'
echo 'export PATH="$HOME/.re/bin:$PATH"'
