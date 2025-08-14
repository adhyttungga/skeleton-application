#!/bin/bash

# Navigate to the project dir
cd "$(dirname "$0")/.."

# Show user service logs
docker logs --follow skeleton-app-user-service-1