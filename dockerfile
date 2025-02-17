
FROM public.ecr.aws/l9s2p5c4/baizeplus/ubuntu:latest

WORKDIR /build
COPY baize /build/baize
COPY /template /build/template
VOLUME ["/build/config"]
EXPOSE 80
ENTRYPOINT ["./baize"]
