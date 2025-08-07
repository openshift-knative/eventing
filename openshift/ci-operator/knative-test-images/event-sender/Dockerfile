# DO NOT EDIT! Generated Dockerfile for test/test_images/event-sender.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.24-openshift-4.20
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./test/test_images/event-sender

FROM $GO_RUNTIME

ARG VERSION=knative-v1.19
RUN microdnf update tzdata -y && microdnf reinstall tzdata -y

COPY --from=builder /usr/bin/main /usr/bin/event-sender
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-eventing-test-test-images-event-sender-rhel8-container" \
      name="openshift-serverless-1/eventing-test-test-images-event-sender-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Eventing Test Test Images Event Sender" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Eventing Test Test Images Event Sender" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Eventing Test Test Images Event Sender" \
      io.k8s.description="Red Hat OpenShift Serverless Eventing Test Test Images Event Sender" \
      io.openshift.tags="test-test-images-event-sender"

ENTRYPOINT ["/usr/bin/event-sender"]
