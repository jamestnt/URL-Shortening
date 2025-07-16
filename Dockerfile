FROM golang:1.24

WORKDIR /app
RUN go env -w GOPROXY=https://proxy.golang.org,direct
RUN go install github.com/air-verse/air@v1.52.3
RUN if [ ! -f go.mod ]; then go mod init github.com/user/project; fi
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd /app/cmd/api
CMD ["air", "-c", ".air.toml"]