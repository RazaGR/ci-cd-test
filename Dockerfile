
FROM golang:buster
RUN \
    apk --update add curl bash nano musl-dev && \
    rm -r /var/cache/apk/*

# The binary is built and downloaded to the current directory by CI.
COPY ./myapp /bin

# Run the program.
ENTRYPOINT /bin/myapp:WORKDIR name
