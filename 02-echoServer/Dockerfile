FROM golang
WORKDIR /app 
COPY go.mod go.sum ./
RUN ls
COPY *.go ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /echo-server
CMD ["/echo-server"]
