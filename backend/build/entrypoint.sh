#!/bin/bash

echo 'Starting entrypoint.sh script...'
if [[ "$DATABASE" = "mysql" ]]; then
    echo 'Waiting for MySql...'

    until nc -z "$DB_HOST" 3306; do
        sleep 5
    done

    echo 'build project'
    go build -o ./cmd/api/main ./cmd/api/main.go

    echo 'MySQL started'
fi

exec "${@}"
