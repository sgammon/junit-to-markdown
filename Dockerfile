FROM golang:1.12-alpine as build
ARG PACKAGE=github.com/pagero/junit-to-markdown
RUN mkdir -p /go/src/${PACKAGE}
WORKDIR /go/src/${PACKAGE}
RUN apk add git make
COPY Makefile go.* ./
RUN make deps
COPY . .
RUN make build

FROM alpine:3.10
ARG PACKAGE=github.com/pagero/junit-to-markdown
ENTRYPOINT ["junit-to-markdown"]
COPY --from=build /go/src/${PACKAGE}/junit-to-markdown /bin/
EXPOSE 8080