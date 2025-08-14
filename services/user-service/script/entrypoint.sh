#!/bin/bash

# Sourcing the .env file
# source ./.env

# Wait for connection
# echo "${MONGO_HOST}:${MONGO_PORT}"
# wait-for "${MONGO_HOST}:${MONGO_PORT}" -- "$@"
wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# Run the app
./user-service