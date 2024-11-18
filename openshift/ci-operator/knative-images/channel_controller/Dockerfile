# DO NOT EDIT! Generated Dockerfile for cmd/in_memory/channel_controller.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/in_memory/channel_controller

FROM $GO_RUNTIME

ARG VERSION=knative-v1.15
RUN microdnf update tzdata -y && microdnf reinstall tzdata -y

COPY --from=builder /usr/bin/main /usr/bin/channel_controller
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-eventing-in-memory-channel-controller-rhel8-container" \
      name="openshift-serverless-1/eventing-in-memory-channel-controller-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Eventing In Memory Channel Controller" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Eventing In Memory Channel Controller" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Eventing In Memory Channel Controller" \
      io.k8s.description="Red Hat OpenShift Serverless Eventing In Memory Channel Controller" \
      io.openshift.tags="in-memory-channel-controller"

ENTRYPOINT ["/usr/bin/channel_controller"]
