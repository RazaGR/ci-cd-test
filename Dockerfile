
FROM golang:buster
# The binary is built and downloaded to the current directory by CI.
COPY ./myapp /bin

# Run the program.
ENTRYPOINT /bin/myapp
