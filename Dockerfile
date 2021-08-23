FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o rabbitmq_mate .


FROM ttbb/rabbitmq:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/rabbitmq/mate

COPY --from=build /opt/sh/compile/pkg/rabbitmq_mate /opt/sh/rabbitmq/mate/rabbitmq_mate

WORKDIR /opt/sh/rabbitmq

CMD ["/usr/local/bin/dumb-init", "bash", "-vx", "/opt/sh/rabbitmq/mate/scripts/start.sh"]