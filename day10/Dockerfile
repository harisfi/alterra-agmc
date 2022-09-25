# syntax=docker/dockerfile:1

############################
# STEP 1 build executable binary
############################
# golang:1.18.6-alpine3.16
FROM golang@sha256:0885b3ffdfbe8e6e8caf07e0b030091cbe1e297c7413ee664fe62d265423bade AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata dumb-init && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /go/bin/day10-agmc-app .

############################
# STEP 2 build a small image
############################
FROM alpine@sha256:0885b3ffdfbe8e6e8caf07e0b030091cbe1e297c7413ee664fe62d265423bade

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /usr/bin/dumb-init /usr/bin/dumb-init

# Copy our static executable
COPY --from=builder /go/bin/day10-agmc-app /go/bin/day10-agmc-app

# Use an unprivileged user.
USER appuser:appuser

# Run the day10-agmc-app binary.
# See https://github.com/gofiber/fiber/issues/1036#issuecomment-841763449
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["sh", "-c", "/go/bin/day10-agmc-app"]