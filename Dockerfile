# trunk-ignore-all(trivy/DS026)
# syntax=docker/dockerfile:1@sha256:ac85f380a63b13dfcefa89046420e1781752bab202122f8f50032edf31be0021

# Build go binary
FROM golang:1.22-alpine@sha256:421bc7f4b90d042c56282bb894451108f8ab886687e1b73abaefad31ab10a14d as build

RUN groupadd -r nonroot && useradd -r -g nonroot nonroot

USER nonroot

WORKDIR /go/src/app

COPY go.mod go.sum /
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" ./cmd/api-server

# Run
FROM gcr.io/distroless/static:nonroot@sha256:e9ac71e2b8e279a8372741b7a0293afda17650d926900233ec3a7b2b7c22a246

COPY --from=build /go/bin/app .

EXPOSE 50051
ENTRYPOINT ["/home/nonroot/app"]