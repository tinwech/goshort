FROM        golang
RUN         mkdir -p /server
WORKDIR     /server
COPY        . .
RUN         go mod download
RUN         go build -o server
ENTRYPOINT  ["./server"]