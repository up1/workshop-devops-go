# Workshop DevOps
* API with [Go](https://golang.org/)
* Docker
  * Build image with [multi-stage build](https://docs.docker.com/develop/develop-images/multistage-build/)
* Docker compose
* CI/CD with [Jenkins](https://www.jenkins.io/)

## Step 1 :: Build and Run API with Docker

Working with Docker command
```
$cd api
$docker image build -t somkiat/demo-api .
$docker container run -d --name api -p 8000:1323 somkiat/demo-api
$docker container ps
```

Working with Docker compose
```
$docker-compose build
$docker-compose up -d
$docker-compose ps
$docker-compose logs --follow
```

## Step 2 :: Design build pipeline with Jenkins
* working with free style job
* working with Jenkinsfile (Pipeline as a Code)
