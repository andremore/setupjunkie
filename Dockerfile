FROM ubuntu:latest

# Avoid prompts
ENV DEBIAN_FRONTEND=noninteractive

# This allows us to get our docker terminal with colors
# In case you want to test what happens if the user doesn't have colors in the terminal
# Just comment this line below
ENV TERM xterm-256color

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
