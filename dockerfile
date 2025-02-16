
FROM 114698169285.dkr.ecr.ap-southeast-1.amazonaws.com/ubuntu:latest

WORKDIR /build
COPY baize /build/baize
COPY /template /build/template
VOLUME ["/build/config"]
EXPOSE 80
ENTRYPOINT ["./baize"]
