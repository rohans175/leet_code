FROM golang:alpine3.20 AS build
WORKDIR /src
COPY app/* /src/
RUN go build -o /bin/test_code test_code.go

FROM golang:alpine3.20
COPY --from=build /bin/test_code /bin/test_code
CMD ["/bin/test_code"]