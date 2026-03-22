# -----------------------------------
# Base image (Debian ARMv7)
# -----------------------------------
FROM arm32v7/debian:bullseye AS builder

# -----------------------------------
# Metadata
# -----------------------------------
LABEL maintainer="paul.fleischmann81@gmail.com"
LABEL description="Cross-compile environment for BeagleBone Black (ARMv7)"

# -----------------------------------
# Tools & dependencies
# -----------------------------------
RUN apt-get update && apt-get install -y \
    build-essential \
    gcc-arm-linux-gnueabihf \
    g++-arm-linux-gnueabihf \
    cmake \
    git \
    wget \
    curl \
    unzip \
    python3 \
    python3-pip \
    ca-certificates \
    sudo \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

# -----------------------------------
# Create non-root user (optional, safer)
# -----------------------------------
RUN useradd -ms /bin/bash bbbuser && \
    echo "bbbuser ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

USER bbbuser
WORKDIR /home/bbbuser

# -----------------------------------
# Copy project files into container
# -----------------------------------
COPY . /home/bbbuser/project

WORKDIR /home/bbbuser/project

# -----------------------------------
# Optional: Entrypoint for debugging / runtime IP check
# -----------------------------------
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

# -----------------------------------
# Default command (override in CI)
# -----------------------------------
CMD ["/bin/bash"]