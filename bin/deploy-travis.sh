#!/usr/bin/env bash
mkdir -p dist
docker login -u $DOCKER_USER -p $DOCKER_PASSWORD $DOCKER_HOSTNAME
export REPO=$DOCKER_USER/educonn-mail
export TAG=`if [ "$TRAVIS_BRANCH" == "master" ]; then echo "latest"; else echo $TRAVIS_TAG ; fi`
echo $REPO:$TAG
docker build -f Dockerfile -t $REPO:$TAG .
docker push $REPO
