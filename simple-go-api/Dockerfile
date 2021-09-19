FROM registry.redhat.io/rhel8/go-toolset:1.15.14-3 AS build_base

USER root

WORKDIR /tmp/go-env

ADD main.go .

RUN go mod init example/webserver && go mod tidy && go build && chgrp root webserver && chmod g+x webserver

FROM registry.redhat.io/ubi8-minimal:8.4-208

COPY --from=build_base /tmp/go-env/webserver /app/webserver

EXPOSE 8000

CMD ["/app/webserver"]
