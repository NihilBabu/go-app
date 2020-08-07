
FROM golang:1.14.4

# Secure against running as root

RUN groupadd --gid 5000 newuser \
    && useradd --home-dir /home/newuser --create-home --uid 5000 \
        --gid 5000 --shell /bin/sh --skel /dev/null newuser
#RUN adduser --disabled-password  -u 10000 nonroot
#RUN mkdir /micro/ && chown nonroot /micro/
#USER nonroot

WORKDIR /micro/
ADD . /micro/

# Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 go build -o /micro .

# final stage
#FROM alpine:3.8
#
## Secure against running as root
#RUN adduser -D -u 10000 nonroot
#USER nonroot
#
#WORKDIR /
#COPY --from=build-env /micro/certs/docker.localhost.* /
#COPY --from=build-env /micro/gcuk /
#
#EXPOSE 8080
#
#CMD ["/gcuk"]