FROM 114698169285.dkr.ecr.ap-southeast-1.amazonaws.com/golang:latest as builder


# 移动到工作目录：/build
WORKDIR /build

COPY . .
RUN go mod download


RUN cd app/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../KnitLogin

FROM 114698169285.dkr.ecr.ap-southeast-1.amazonaws.com/ubuntu:latest

WORKDIR /build/knit
COPY --from=builder /build/KnitLogin .
COPY --from=builder /build/i18n ./i18n
VOLUME ["/build/knit/config"]
EXPOSE 80
ENTRYPOINT ["./KnitLogin"]
