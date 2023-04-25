# Do not edit! This file was generated via Makefile
FROM registry.ci.openshift.org/openshift/release:golang-1.18 as builder

COPY . .
RUN make install test-install

FROM registry.access.redhat.com/ubi8/ubi-minimal

RUN microdnf install tzdata

USER 65532

COPY ./cmd/in_memory/channel_dispatcher/kodata /var/run/ko
COPY --from=builder /go/bin/channel_dispatcher /usr/bin/channel_dispatcher
ENTRYPOINT ["/usr/bin/channel_dispatcher"]
