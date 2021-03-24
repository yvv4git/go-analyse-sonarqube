#!/bin/bash
SONARQUBE_URL=192.168.0.106:9777
MY_REPO=$(pwd)

echo "HOST ADDR: "$SONARQUBE_URL;
echo "PATH: "$MY_REPO;

docker run \
    --rm \
    --network host \
    -e SONAR_HOST_URL="http://${SONARQUBE_URL}" \
    -e SONAR_LOGIN="f6c3f29215fd27924ba42bd57553f182588e4324" \
    -v "${MY_REPO}:/usr/src" \
    sonarsource/sonar-scanner-cli