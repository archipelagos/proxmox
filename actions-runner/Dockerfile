FROM ubuntu:latest

# INFO: Download mandatory software.
RUN apt update && apt install -y adduser

# INFO: Set workdir.
WORKDIR /home/ubuntu

# INFO: Download mandatory software.
RUN apt update && apt install -y curl

# INFO: Download the latest runner package.
RUN curl -o actions-runner-linux-x64-2.332.0.tar.gz -L https://github.com/actions/runner/releases/download/v2.332.0/actions-runner-linux-x64-2.332.0.tar.gz

# INFO: Validate the hash.
RUN echo "f2094522a6b9afeab07ffb586d1eb3f190b6457074282796c497ce7dce9e0f2a  actions-runner-linux-x64-2.332.0.tar.gz" | sha256sum -c

# INFO: Extract the installer.
RUN mkdir -p actions-runner
RUN tar xzf actions-runner-linux-x64-*.tar.gz -C actions-runner
RUN rm -rf actions-runner-linux-x64-*.tar.gz

# INFO: Set proper ownership.
RUN chown -R ubuntu:ubuntu /home/ubuntu

# INFO: Install all needed dependencies.
RUN cd actions-runner && ./bin/installdependencies.sh

# INFO: Change user to standard one.
USER 1000

# INFO: Set workdir.
WORKDIR /home/ubuntu/actions-runner

# INFO: Copy helper script.
COPY entrypoint.sh entrypoint.sh

# INFO: Set entrypoint.
ENTRYPOINT [ "bash", "-c", "/home/ubuntu/actions-runner/entrypoint.sh" ]
