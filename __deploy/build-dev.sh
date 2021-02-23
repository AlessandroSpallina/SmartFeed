cd ../identity-node/src
docker build --rm -t identity-node:v0 . -f Dockerfile-dev

cd ../../feed-producer
docker build --rm -t feed-producer:v0 . -f Dockerfile-dev

cd ../feed-provider
docker build --rm -t feed-provider:v0 . -f Dockerfile-dev

# grafana dashboard stuff
# see https://stackoverflow.com/questions/34031397/running-docker-on-ubuntu-mounted-host-volume-is-not-writable-from-container
#chmod go+rw ../../grafana
