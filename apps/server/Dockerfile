# One static binary on nothing, per ADR-0014. CGO_ENABLED=0 is what makes both
# this image and the eventual self-hosted rung simple.
FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH:-amd64} \
	go build -trimpath -ldflags="-s -w" -o /out/nah-server ./cmd/server
# Carried over so the runtime image has a data directory it owns.
RUN mkdir -p /out/data && chown -R 65532:65532 /out/data

FROM scratch
COPY --from=build /out/nah-server /nah-server
COPY --from=build --chown=65532:65532 /out/data /data
USER 65532:65532
ENV NAH_ADDR=":8080" NAH_DATA_DIR="/data"
EXPOSE 8080
VOLUME ["/data"]
# Health is checked from outside on /healthz; there is no shell in here to do it.
ENTRYPOINT ["/nah-server"]
