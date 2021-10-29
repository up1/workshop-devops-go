#!/bin/bash
echo "Starting db..."
docker-compose down
docker-compose up -d db

STATUS="starting"
while [ "$STATUS" != "healthy" ]
do
    STATUS=$(docker inspect --format {{.State.Health.Status}} db)
    echo "API state = $STATUS"
    sleep 5
done

echo "Starting api..."
docker-compose up -d api
echo "Finish.."
