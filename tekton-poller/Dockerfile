FROM alpine:3.22

ARG KUBECTL_VERSION=v1.34.1

RUN apk add --no-cache \
    bash \
    ca-certificates \
    curl \
    git

RUN curl -fsSL \
    "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl" \
    -o /usr/local/bin/kubectl \
    && chmod +x /usr/local/bin/kubectl

RUN addgroup -S poller \
    && adduser -S -G poller poller

USER poller

WORKDIR /home/poller

CMD ["/bin/bash"]