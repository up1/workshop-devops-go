#!/bin/bash
echo "Starting api..."

docker-compose -f docker-compose-deploy.yml down
docker-compose -f docker-compose-deploy.yml up -d

STATUS="starting"
while [ "$STATUS" != "healthy" ]
do
    STATUS=$(docker inspect --format {{.State.Health.Status}} api)
    echo "API state = $STATUS"
    sleep 5
done