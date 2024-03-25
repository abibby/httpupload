# syntax=docker/dockerfile:1

FROM golang:1.22

# Set destination for COPY
WORKDIR /app


# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-http-upload

# http server port
EXPOSE 80

ENV PORT=80
ENV ROOT_DIR=/files

# Run
CMD ["/docker-http-upload"]