# Go Api Server
A few things to know:
* **[Github Actions Workflows](https://github.com/polibenitez/go-api-server/tree/main/.github/workflows)** - Pre-configured Github Actions to run automated builds and publish image to Github Packages
* **[Dockerfile](https://github.com/polibenitez/go-api-server/blob/main/Dockerfile)** - Dockerfile to generate docker builds.
* **[docker-compose](https://github.com/polibenitez/go-api-server/blob/main/docker-compose.yml)** - Docker compose script to start service with database.
* **[.env file for configuration](#environment)** - Environment variables
* **[Terraform files](https://github.com/polibenitez/go-api-server/tree/main/infrastructure)** - Terraform file where is the infrastructure as a code
## Development

### Infrastructure
Inside infrstructure folder there is a terraform file. This file contains all necessary resources in AWS to deploy this app as a container in ECS.

### Automate deployment
To automate the deployment, there is a github action pipeline in github/workflow folder. This pipeline contain all necesary steps to build and deploy the app in ECS.

### Start dev server
Starting the dev server also starts MongoDB as a service in a docker container using the compose script at `docker-compose.yml`.

Running the above commands results in 
* 🌏 **API Server** running at `http://localhost:9999`
* ⛁ **MongoDB** running at `mongodb://localhost:27017/db`

## Packaging and Runing
#### 1. Build and run without Docker

```
$ go install
$ go run .
```
#### 3. Run with docker

```
$ docker build -t recipes-api .
$ docker-compose up -d
```
---

## Environment

| Var Name  | Type  | Default | Description  |
|---|---|---|---|
|  PORT | number  | `9999` | Port to run the API server on |
|  MONGO_URL | string  | `mongodb://localhost:27017/db` | URL for MongoDB |