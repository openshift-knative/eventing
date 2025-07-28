# DO NOT EDIT! Generated Dockerfile for cmd/broker/filter.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.24-openshift-4.20
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/broker/filter

FROM $GO_RUNTIME

ARG VERSION=knative-v1.19
RUN microdnf update tzdata -y && microdnf reinstall tzdata -y

COPY --from=builder /usr/bin/main /usr/bin/filter
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-eventing-broker-filter-rhel8-container" \
      name="openshift-serverless-1/eventing-broker-filter-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Eventing Broker Filter" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Eventing Broker Filter" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Eventing Broker Filter" \
      io.k8s.description="Red Hat OpenShift Serverless Eventing Broker Filter" \
      io.openshift.tags="broker-filter"

ENTRYPOINT ["/usr/bin/filter"]
