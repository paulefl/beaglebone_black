# =============================================
# BBB Cross-Compile & Test Builder Container
# Supports C, Rust, Go, Python
# =============================================
FROM debian:12-slim

# ----------------------------
# Install basic packages
# ----------------------------
RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    curl \
    git \
    ca-certificates \
    pkg-config \
    python3 python3-pip python3-setuptools \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# ----------------------------
# Install C Cross-Compiler for ARMv7
# ----------------------------
RUN apt-get update && apt-get install -y \
    gcc-arm-linux-gnueabihf \
    g++-arm-linux-gnueabihf \
    && rm -rf /var/lib/apt/lists/*

# ----------------------------
# Install Rust + ARM Target
# ----------------------------
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"

# Add ARMv7 target
RUN rustup target add armv7-unknown-linux-gnueabihf

# Create Cargo config for cross-linker
RUN mkdir -p /root/.cargo && \
    echo '[target.armv7-unknown-linux-gnueabihf]\nlinker = "arm-linux-gnueabihf-gcc"' > /root/.cargo/config.toml

# ----------------------------
# Install Go + Cross-Compile support
# ----------------------------
RUN curl -OL https://go.dev/dl/go1.21.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.2.linux-amd64.tar.gz && \
    rm go1.21.2.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# ----------------------------
# Set working directory
# ----------------------------
WORKDIR /src

# ----------------------------
# Entrypoint to allow arbitrary commands
# ----------------------------
ENTRYPOINT ["/bin/bash"]