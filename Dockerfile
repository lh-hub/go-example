FROM golang:1.17 AS build
WORKDIR /httpserver/
COPY demo01.go .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver demo01.go

FROM busybox
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 8360
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]
