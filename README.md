# Development

## Running integration tests

Run development containers:
```
docker-compose -f docker-compose-dev.yml up -d
```

Execute tests:
```
ND_IP_PORT=YOUR_DOCKER_MACHINE_IP go test -v -tags integration
```

## Testing locally

Build project with `nildev`, it will produce docker image:
```
nildev build {{.ProjectPath}}
```

Create `local.env` file in root dir with content:
```
ND_DATABASE_NAME={{.TableName}}
ND_MONGODB_URL=mongodb://YOUR_DOCKER_MACHINE_IP:27017/nildev
ND_ENV=dev
```

Run containers:
```
docker-compose -f docker-compose.yml up -d
```

Execute HTTP request:
```
curl -X POST http://YOUR_DOCKER_MACHINE_IP/api/v1/custom-register/github?userName=nildev -v
```

## Project Details

### Release Notes

See the [releases tab](https://github.com/nildev/{{.ApiName}}/releases) for more information on each release.

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches and contacting developers via IRC and mailing lists.

### License

Project is released under the MIT license. See the [LICENSE](LICENSE) file for details.

Specific components of project use code derivative from software distributed under other licenses; in those cases the appropriate licenses are stipulated alongside the code.