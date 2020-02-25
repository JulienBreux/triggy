#
# Builder
#
FROM golang:1.13.6-alpine3.11 AS builder

# Define application directory
WORKDIR /app

# Copy app's files
COPY go.mod go.sum cmd/triggy/triggy.go ./
COPY internal internal
COPY cmd cmd
COPY Taskfile.yml Taskfile.yml

# Install deps
RUN apk --no-cache add git gcc && \
    # Get go-task
    TASK_VERSION=2.8.0 && \
    wget -q https://github.com/go-task/task/releases/download/v${TASK_VERSION}/task_linux_amd64.tar.gz && \
    tar -xzf task_linux_amd64.tar.gz task && \
    mv task /usr/local/bin/ && \
    chmod ugo+x /usr/local/bin/task && \
    rm task_linux_amd64.tar.gz && \
    # Build app
    task app:build

#
# App
#
FROM alpine:3.11

# Labels
LABEL Maintainer="Julien BREUX <julien.breux@gmail.com>"

# Copy app
COPY --from=builder /app/triggy /bin/triggy

# Install required deps
RUN apk add --update ca-certificates && \
    # Install temp deps
    # apk add --update -t deps ... && \
    # Purge deps
    # apk del --purge deps && \
    rm /var/cache/apk/*

# Define entry point
ENTRYPOINT [ "/bin/triggy" ]
