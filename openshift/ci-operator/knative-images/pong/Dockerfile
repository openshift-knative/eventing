# Do not edit! This file was generated via Makefile
FROM registry.ci.openshift.org/openshift/release:golang-1.18 as builder

COPY . .
RUN make install test-install

FROM registry.access.redhat.com/ubi8/ubi-minimal

RUN microdnf install tzdata

USER 65532

COPY ./cmd/pong/kodata /var/run/ko
COPY --from=builder /go/bin/pong /usr/bin/pong
ENTRYPOINT ["/usr/bin/pong"]
