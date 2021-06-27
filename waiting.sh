#!/bin/bash
echo "Starting api..."
STATUS="starting"
while [ "$STATUS" != "healthy" ]
do
    STATUS=$(docker inspect --format {{.State.Health.Status}} api)
    echo "API state = $STATUS"
    sleep 5
done