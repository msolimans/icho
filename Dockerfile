FROM golang:alpine AS buildstg
ENV GOPATH /wd
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ADD . /go
RUN cd /go && go build -mod=vendor -ldflags="-w -s" -o /icho
FROM scratch
COPY --from=buildstg /icho /icho
#RUN ls /
ENTRYPOINT ["/icho"]
#https://github.com/chemidy/smallest-secured-golang-docker-image/blob/master/go_module/Dockerfile
#ADD "https://curl.haxx.se/ca/cacert.pem" "/etc/ssl/certs/ca-certificates.crt"
#ADD "./pkg/linux_amd64/http-echo" "/"
#ENTRYPOINT ["/http-echo"]