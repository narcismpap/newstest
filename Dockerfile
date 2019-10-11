# Package: com.github.narcismpap.news
# Dockerfile
#
# Author: Narcis M. Pap on 22/06/2019

ARG GO_VERSION=1.12
FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN apk add --no-cache ca-certificates git
ENV CGO_ENABLED=0

WORKDIR /src
COPY ./server ./
COPY ./config.json ./config.json

RUN go get -d -v ./...
RUN go build \
    -installsuffix 'static' \
    -o /news .


FROM scratch AS final
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /news /news

EXPOSE 9000
USER nobody:nobody

COPY ./config.json ./config.json
ENTRYPOINT ["/news"]
