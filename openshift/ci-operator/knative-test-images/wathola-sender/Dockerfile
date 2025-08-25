# DO NOT EDIT! Generated Dockerfile for test/test_images/wathola-sender.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.23-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./test/test_images/wathola-sender

FROM $GO_RUNTIME

ARG VERSION=knative-v1.15
RUN microdnf update tzdata -y && microdnf reinstall tzdata -y

COPY --from=builder /usr/bin/main /usr/bin/wathola-sender
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-eventing-test-test-images-wathola-sender-rhel8-container" \
      name="openshift-serverless-1/eventing-test-test-images-wathola-sender-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Eventing Test Test Images Wathola Sender" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Eventing Test Test Images Wathola Sender" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Eventing Test Test Images Wathola Sender" \
      io.k8s.description="Red Hat OpenShift Serverless Eventing Test Test Images Wathola Sender" \
      io.openshift.tags="test-test-images-wathola-sender"

ENTRYPOINT ["/usr/bin/wathola-sender"]
