FROM golang:1.22-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/dr-demo-telemetry-collector .

FROM alpine:3.19
COPY --from=build /out/dr-demo-telemetry-collector /usr/local/bin/dr-demo-telemetry-collector
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/dr-demo-telemetry-collector"]
