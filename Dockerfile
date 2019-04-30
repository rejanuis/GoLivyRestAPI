###############################
# STEP 1 build executable binary
###############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Add folder swaggerui for running with swagger-ui
ADD ./swaggerui /swaggerui
# Set work directory before compile
COPY . /home/GoLivyRestAPI/
WORKDIR /home/GoLivyRestAPI/
# Fetch dependencies.
# Using go get.
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/rs/cors
RUN go get -u github.com/tidwall/gjson
RUN go get -u github.com/spf13/viper
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /home/GoLivyRestAPI/golivyrest
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /home/GoLivyRestAPI/golivyrest /home/
COPY --from=builder /swaggerui/ /home/swaggerui/
# Set Directory for smallest image
WORKDIR /home
# COPY  controller /controller/
# COPY  model /model/
# COPY  main /bin/main
# CMD ["./main"]
# # Run the binary.
ENTRYPOINT ["/home/golivyrest"]
# ENTRYPOINT ["/home/demo"]