FROM alpine:3.6
WORKDIR /app
# Now just add the binary
COPY go2hal /app/
RUN apk add --no-cache openssh-client
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ENTRYPOINT ["/app/go2hal"]
EXPOSE 8000 8080