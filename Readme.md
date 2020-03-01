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

### Test in browser

Open you browser and point to http://localhost:8888/echo, you should see `{response: "hello world"}`

If you want to override  response you can do that through query string, e.g. http://localhost:8888/echo?text=foo

If you want to override response through environment variable

```shell script
$ docker run -ti --rm \
   -e ECHO_TEXT="foo response" \
   -p 8888:8888 -p 9999:9999 \
   msolimans/icho
$ curl http://localhost:8888/echo
{"response":"foo response"}
```     

If you want to try gRPC, open `proto` file under `icho` directory in any of the gRPC client tools and use port `9999`
  
Here's gRPC Client, https://github.com/uw-labs/bloomrpc