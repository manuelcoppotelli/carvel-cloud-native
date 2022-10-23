FROM golang:1.18 as build
COPY . .
RUN CGO_ENABLED=0 go build -o app app.go

FROM gcr.io/distroless/static-debian11
COPY --from=build /go/app /
CMD ["/app"]
