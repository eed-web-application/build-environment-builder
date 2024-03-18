# Use the specified base image
FROM golang:latest

# Install Node.js
RUN apt update \
    && apt-get install -y nodejs npm default-jdk \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Continue with your container setup
# ...
RUN npm install @openapitools/openapi-generator-cli -g