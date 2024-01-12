FROM scratch
# COPY app.run /server
# COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# COPY nsswitch.conf /etc/nsswitch.conf
COPY docs/swagger.json docs/swagger.json
COPY logs logs
CMD [ "/server" ]

# Dockerfile for the Go service
FROM golang:latest

WORKDIR /project

# Copy the project files into the container
COPY . .

# Set the environment variables
ENV TZ=Asia/Jakarta
ENV PORT=8000

# Expose the port
EXPOSE 8000

# Command to run when the container starts
CMD ["tail", "-f", "/etc/hosts"]