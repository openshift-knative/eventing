# DO NOT EDIT! Generated Dockerfile for cmd/websocketsource.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.25-openshift-4.21
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/websocketsource

FROM $GO_RUNTIME

ARG VERSION=knative-v1.16
RUN microdnf update tzdata -y && microdnf reinstall tzdata -y

COPY --from=builder /usr/bin/main /usr/bin/websocketsource
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-eventing-websocketsource-rhel8-container" \
      name="openshift-serverless-1/eventing-websocketsource-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Eventing Websocketsource" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Eventing Websocketsource" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Eventing Websocketsource" \
      io.k8s.description="Red Hat OpenShift Serverless Eventing Websocketsource" \
      io.openshift.tags="websocketsource" \
      vendor="Red Hat, Inc." \
      release=$VERSION \
      cpe="cpe:/a:redhat:openshift_serverless:1.36::el8"
      
ENTRYPOINT ["/usr/bin/websocketsource"]
