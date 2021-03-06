# Icho

Simple echo server that echos messages, its supports both gRPC and Http

### Build locally 

```shell script
$ git clone https://github.com/msolimans/icho 
$ cd icho 
$ go build -mod=vendor
```

### Docker 

```shell script
$ docker pull msolimans/icho 
$ docker run -ti --rm -p 8888:8888 -p 9999:9999 msolimans/icho
```  

### Testing 

Open you browser and point to http://localhost:8888/echo, you should see `{response: "hello world"}`

If you want to override  response you can do that through query string, something like

http://localhost:8888/echo?text=foo

If you want to echo based on http `POST`, try the following:

```shell script
$ curl  --request POST --data-raw '"foo message"' \
    http://localhost:8888/echo
  {"response":"foo message"}
``` 

If you want to override response through environment variable

```shell script
$ docker run -ti --rm \
   -e ECHO_TEXT="foo response" \
   -p 8888:8888 -p 9999:9999 \
   msolimans/icho
$ curl http://localhost:8888/echo
{"response":"foo response"}
```     

If you want to override ports, you can pass `http` and `grpc` ports like below 
```shell script
$ docker run --rm -ti \
-p 7777:7777 \
-p 6666:6666  \
msolimans/icho -http=6666 -grpc=7777

Starting REST service :6666 
Starting gRPC service :7777
```
 
 
If you want to try gRPC, open `proto` file under `icho` directory in any of the gRPC client tools and use port `9999`
  
Here's gRPC Client, https://github.com/uw-labs/bloomrpc

### For Kubernetes: 

Create deployment resource:
```shell script
$ touch icho.yml 
$ vim icho.yml 
```  

Add the following: 

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    run: icho
  name: icho
spec:
  replicas: 1
  selector:
    matchLabels:
      run: icho
  template:
    metadata:
      labels:
        run: icho
    spec:
      containers:
      - image: msolimans/icho
        name: icho
        ports:
        - containerPort: 8888
```

Create deployment: 

```shell script
$ kubectl create -f icho.yml 
```


Expose deployment as a `NodePort` service: 

```shell script
$ kubectl expose deploy \
  --name ichosvc \
  --target-port=8888 --port=8888 \
  --type=NodePort \
  icho 
```

open your browser http://localhost:8888/echo 

If you're using `minikube`, open it using:

```shell script
$ minikube service ichosvc 
```

### docker-compose

```shell script
$ git clone https://github.com/msolimans/icho
$ cd icho 
$ docker-compose up 
```