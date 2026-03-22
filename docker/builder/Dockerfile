FROM ubuntu:22.04
RUN apt-get update && apt-get install -y build-essential \
    gcc-arm-linux-gnueabihf rustc cargo golang-go python3 \
    python3-pip make git \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /src
CMD ["/bin/bash"]