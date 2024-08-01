FROM golang:alpine3.20 AS build
WORKDIR /src
COPY code/* ./
RUN go build -o /bin/main main.go

FROM golang:alpine3.20
COPY --from=build /bin/main /bin/main
CMD [ "/bin/main" ]
