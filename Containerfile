FROM registry.access.redhat.com/ubi9/ubi-minimal
WORKDIR /
COPY _output/kube-rbac-proxy /usr/local/bin/kube-rbac-proxy
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/usr/local/bin/kube-rbac-proxy"]
