#!/usr/bin/env bash

ROOT_DIR="$(cd "$(dirname "$0")/.." &> /dev/null; pwd -P)"
MIGRATIONS_DIR="migrations"

. "${ROOT_DIR}/.env"

cd "${ROOT_DIR}/${MIGRATIONS_DIR}"
for f in $(ls .); do
	echo "Exec migration $f..."
	mysql -u"$DB_USER" "$DB_NAME" < "$f"
done 