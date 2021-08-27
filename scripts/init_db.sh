#! /bin/sh
set -x
set -eo pipefail

if ! [ -x "$(command -v psql)" ]; then
	echo >&2 "Error: `psql` is not installed"
	exit 1
fi

if ! [ -x "$(command -v migrate)" ]; then
	echo >&2 "Error: `migrate` is not installed"
	echo >&2 "See: "
	echo >&2 "	https://github.com/golang-migrate/migrate/tree/master/cmd/migrate "
	echo >&2 "and follow the instructions to install it."
	exit 1
fi

DB_USER=${POSTGRES_USER:=postgres}
DB_PASSWORD=${POSTGRES_PASSWORD:=password}
DB_NAME=${POSTGRES_DB:=market_data}
DP_PORT=${POSTGRES_PORT:=5432}

if [[ -z "${SKIP_DOCKER}" ]]
then
	docker-compose up -d
fi


until psql -h "localhost" -U "${DB_USER}" -p "${DP_PORT}" -d "${DB_NAME}" -c '\q'; do
	>&2 echo "Postgres is still unvailable - sleeping"
	sleep 1
done

>&2 echo "Postgres is up and running on port ${DB_PORT}!"
export DATABASE_URL=postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}

migrate create -ext sql -dir db/migrations -seq create_users_table

migrate -path db/migrations -database "${DATABASE_URL}?sslmode=disable" --verbose up

>&2 echo "Postgres has been migrated, ready to go."
