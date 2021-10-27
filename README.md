# Workshop DevOps
* API with [Go](https://golang.org/)
* Docker
  * Build image with [multi-stage build](https://docs.docker.com/develop/develop-images/multistage-build/)
  * [Best practices for writing Dockerfiles](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
* Docker compose
* CI/CD with [Jenkins](https://www.jenkins.io/)
* Static Code Analysis
  * [SonarQube](https://www.sonarqube.org/)
  * [Sonar Scanner](https://docs.sonarqube.org/latest/analysis/scan/sonarscanner/)

sonar-scanner.properties
```
sonar.host.url=http://128.199.199.84:9000
sonar.login=xxx
sonar.password=xxx
```

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

Deploy with shell script
```
$sh deploy.sh
$docker-compose ps
$docker-compose logs --follow
```

URL for testing
* http://localhost:8000/accounts/1
* http://localhost:8000/accounts/2
* http://localhost:8000/accounts/3


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

## Step 3 :: Deploy with docker swarm

Create Docker Swarm
```
$docker swarm init
$docker node ls
```

See join-token command
```
$docker swarm join-token worker
$docker swarm join-token manager
```

Deploy stack with docker-compose file (v3+)
```
$docker stack deploy --compose-file docker-compose-deploy.yml demostack
$docker stack services demostack
```

Scaling service
```
docker service ls
$docker service scale <service name>=<number of replica(s)>
$docker service ls
$docker service ps <service name>
```

Example scale to 5
```
$docker service scale demostack_api=5
```

demostack_api scaled to 5
overall progress: 5 out of 5 tasks
1/5: running   [==================================================>]
2/5: running   [==================================================>]
3/5: running   [==================================================>]
4/5: running   [==================================================>]
5/5: running   [==================================================>]
verify: Service converged

Testing with curl and see result
```
$curl http://localhost:8000/
```

Delete services
```
$docker service rm <service name>
```

Delete stack
```
$docker stack rm demostack
```

Leave node from Swarm
```
$docker swarm leave --force
```

## Step 4 :: Deploy with Kubernetes
```
$kubectl get node
$kubectl apply -f k8s/
$kubectl get all
```
