## Build and Run with Docker
```
$docker image build -t somkiat/demo-api .
$docker container run -d --name api -p 8000:1323 somkiat/demo-api
$docker container ps

CONTAINER ID   IMAGE              COMMAND   CREATED         STATUS         PORTS                                       NAMES
1f573d40e1d0   somkiat/demo-api   "./api"   4 seconds ago   Up 2 seconds   0.0.0.0:8000->1323/tcp, :::8000->1323/tcp   api

$docker logs --follow api

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.3.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```

Access to api with http://localhost:8000/