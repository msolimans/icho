# Icho

Simple echo server that echos messages, its supports both gRPC and Http

### Build locally 

```
$ git clone https://github.com/msolimans/icho 
$ cd icho 
$ go build -mod=vendor
```

### Docker 

```
$ docker pull msolimans/icho 
$ docker run -ti --rm -p 8888:8888 -p 9999:9999 msolimans/icho
```  

Open http://localhost:8888 in the browser   

If you want to try gRPC, open `proto` file under `icho` directory in any of the gRPC client tools and use port `9999`
  
Here's gRPC Client, https://github.com/uw-labs/bloomrpc