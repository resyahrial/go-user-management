# app
export APP_NAME=""
export APP_DEBUG_MODE=true
export APP_HOST="localhost"
export APP_PORT="8080"

# db
export DB_HOST=localhost
export DB_USERNAME=postgres
export DB_PASSWORD=postgres
export DB_PORT=5432
export DB_NAME=postgres
export DATABASE_URL=postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable

# hasher
export HASHER_COST=0

# auth
export ACCESS_TIME_DURATION=0
export ACCESS_SECRET_KEY=""

# seed
export SEED_ADMIN_EMAIL=""
export SEED_ADMIN_PASSWORD=""
