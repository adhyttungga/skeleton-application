#!/bin/bash

# Navigate to the project dir
cd "$(dirname "$0")/.."

# Stop and remove docker container
docker-compose down