# Multi-stage build
  FROM node:20 AS builder

  # Install Go
  RUN curl -L https://go.dev/dl/go1.22.0.linux-amd64.tar.gz | tar -xz -C /usr/local
  ENV PATH="/usr/local/go/bin:${PATH}"

  # Install build dependencies
  RUN apt-get update && apt-get install -y git build-essential

  # Increase Node.js memory limit
  ENV NODE_OPTIONS="--max_old_space_size=4096"

  WORKDIR /app

  # Copy source code
  COPY . .

  # Build using make
  RUN make cli-only

  # Final stage - Ubuntu runtime
  FROM ubuntu:latest

  RUN apt-get update && \
      apt-get install -y ca-certificates && \
      rm -rf /var/lib/apt/lists/*

  # Copy the built binary
  COPY --from=builder /app/rill /usr/local/bin/rill
  RUN chmod +x /usr/local/bin/rill

  RUN groupadd -g 1001 rill && \
      useradd -m -u 1001 -s /bin/sh -g rill rill

  USER rill
  WORKDIR /home/rill

  ENTRYPOINT ["rill"]
  CMD ["start"]
