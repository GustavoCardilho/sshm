#!/usr/bin/env bash
# Instala o sshm: compila o programa, copia para uma pasta no PATH
# e adiciona essa pasta ao ~/.zshrc se ainda não estiver lá.
#
# Uso:
#   ./install.sh                 instala em ~/.local/bin
#   ./install.sh /outra/pasta    instala em outra pasta

set -euo pipefail

BIN_NAME="sshm"
DEST="${1:-$HOME/.local/bin}"
RC_FILE="$HOME/.zshrc"
EXPORT_LINE="export PATH=\"$DEST:\$PATH\""

# Roda sempre a partir da pasta do script, não importa de onde ele foi chamado
cd "$(dirname "${BASH_SOURCE[0]}")"

if ! command -v go >/dev/null 2>&1; then
  echo "Go não encontrado. Instale o Go primeiro: https://go.dev/dl/"
  exit 1
fi

mkdir -p "$DEST"

echo "Compilando..."
go build -o "$DEST/$BIN_NAME" .
echo "Binário instalado em $DEST/$BIN_NAME"

# Só mexe no .zshrc se a pasta ainda não estiver no PATH
# e se a linha ainda não tiver sido adicionada antes
case ":$PATH:" in
  *":$DEST:"*)
    echo "A pasta $DEST já está no seu PATH."
    ;;
  *)
    touch "$RC_FILE"
    if grep -qF "$EXPORT_LINE" "$RC_FILE"; then
      echo "A linha do PATH já existe em $RC_FILE."
    else
      printf '\n# sshm\n%s\n' "$EXPORT_LINE" >> "$RC_FILE"
      echo "Linha do PATH adicionada em $RC_FILE."
    fi
    echo "Rode 'source $RC_FILE' ou abra um terminal novo para usar o comando '$BIN_NAME'."
    ;;
esac

echo "Pronto."