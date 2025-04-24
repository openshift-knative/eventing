# DO NOT EDIT! Generated Dockerfile for cmd/mtping.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.23-openshift-4.19
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/mtping

FROM $GO_RUNTIME

ARG VERSION=knative-v1.18
RUN microdnf update tzdata -y && microdnf reinstall tzdata -y

COPY --from=builder /usr/bin/main /usr/bin/mtping
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-eventing-mtping-rhel8-container" \
      name="openshift-serverless-1/eventing-mtping-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Eventing Mtping" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Eventing Mtping" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Eventing Mtping" \
      io.k8s.description="Red Hat OpenShift Serverless Eventing Mtping" \
      io.openshift.tags="mtping"

ENTRYPOINT ["/usr/bin/mtping"]
