# Workshop DevOps
* API with [Go](https://golang.org/)
* Docker
  * Build image with [multi-stage build](https://docs.docker.com/develop/develop-images/multistage-build/)
* Docker compose
* CI/CD with [Jenkins](https://www.jenkins.io/)

## Step 0 :: Preparaion
```
$git clone https://github.com/up1/workshop-devops-go.git
cd workshop-devops-go
```

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

Deploy with Docker compose
```
$docker-compose -f docker-compose-deploy.yml build
$docker-compose -f docker-compose-deploy.yml up -d
$docker-compose -f docker-compose-deploy.yml ps
$docker-compose -f docker-compose-deploy.yml logs --follow
```

## Step 2 :: Design build pipeline with Jenkins
* working with free style job
* working with [Jenkinsfile](https://www.jenkins.io/doc/book/pipeline/jenkinsfile/) (Pipeline as a Code)

Design youe pipeline
* Pull code
* Build
* Testing
* Build Docker image
* Push Docker image to Registry server
* Deploy to target server
  * Docker
  * Kubernetes cluster 

List of Jenkins's plugins
* [Publish Over SSH](https://plugins.jenkins.io/publish-over-ssh/)
* [Kubernetes Continuous Deploy Plugin 1.0.0](https://www.jenkins.io/doc/pipeline/steps/kubernetes-cd/)
