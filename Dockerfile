FROM golang:1.20.6

RUN go version
ENV GOPATH=/

COPY ./ ./

# DEPENDENCIES
RUN go mod download
# BUILD
RUN go build -o go-start ./cmd
# START
CMD ["./go-start"]
