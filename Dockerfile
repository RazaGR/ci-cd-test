
FROM alpine:3.18.4
RUN \
    apk --update add curl bash nano && \
    rm -r /var/cache/apk/*

# The binary is built and downloaded to the current directory by  CI.
COPY ./myapp /bin

# Run the program.
ENTRYPOINT /bin/myapp
