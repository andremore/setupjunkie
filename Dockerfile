FROM ubuntu:latest

# Avoid prompts
ENV DEBIAN_FRONTEND=noninteractive

# Update and install basic utilities
RUN apt-get update && apt-get install -y \
    sudo \
    curl \
    git \
    vim

# Optionally set a default user (change 'myuser' to whatever you like)
RUN useradd -m myuser && echo "myuser:myuser" | chpasswd && adduser myuser sudo

# Set the user for subsequent commands
USER myuser

# Set working directory
WORKDIR /home/myuser

# Copy your CLI tool binaries or scripts into the container
COPY ./setupjunkie /home/myuser/

# This will keep the container running for testing purposes
CMD ["tail", "-f", "/dev/null"]
