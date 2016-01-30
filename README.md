# Development

## Running integration tests

Create docker machine:
```
docker-machine create --driver virtualbox nildev
```

Setup docker environment:
```
docker-machine start nildev
eval $(docker-machine env nildev)
DM_NILDEV_IP=$(docker-machine ip nildev)
```

Run docker containers:
```
docker run --name local-mongo-nildev -p 27017:27017 -d mongo:3.0.8
```

Execute tests

Provisioning is happening inside `TestMain`
```
DM_NILDEV_IP=$(docker-machine ip nildev) \
ND_MONGODB_URL="mongodb://$DM_NILDEV_IP:27017/nildev" \
ND_DATABASE_NAME="nildev" \
go test -v -tags integration
```

## Testing locally

```
# Build project container, it will produce docker image `docker images`
nildev build
# Get IP 
DM_NILDEV_IP=$(docker-machine ip nildev)
# Run container
docker run -d -p "8080:8080" -e "ND_DATABASE_NAME=nildev" -e "ND_MONGODB_URL=mongodb://$DM_NILDEV_IP:27017/nildev" {{.Org}}/{{.ApiName}}:latest
curl -X POST --data "@register.json" http://$DM_NILDEV_IP:8080/api/v1/Register -v
```

## Project Details

### Release Notes

See the [releases tab](https://github.com/nildev/{{.ApiName}}/releases) for more information on each release.

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches and contacting developers via IRC and mailing lists.

### License

Project is released under the MIT license. See the [LICENSE](LICENSE) file for details.

Specific components of project use code derivative from software distributed under other licenses; in those cases the appropriate licenses are stipulated alongside the code.