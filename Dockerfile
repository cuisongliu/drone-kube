FROM golang:1.11
ADD drone-kube /usr/local/bin/drone-kube
RUN apk --no-cache add curl ca-certificates bash && \
    curl -Lo /usr/local/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && \
    chmod +x /usr/local/bin/* && \
    rm -rf /var/cache/apk/*
ENTRYPOINT ["/usr/local/bin/drone-kube"]
