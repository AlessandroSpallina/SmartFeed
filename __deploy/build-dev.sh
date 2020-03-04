cd ../identity-node/src
docker build --rm -t identity-node:v0 . -f Dockerfile-dev

cd ../feed-producer
docker build --rm -t feed-producer:v0 . -f Dockerfile-dev
