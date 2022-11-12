# dev database, per the machine setup
export DB_HOST=localhost
export DB_USERNAME=postgres
export DB_PASSWORD=postgres
export DB_PORT=5433
export DB_NAME=postgres
export DATABASE_URL=postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable
