FROM ubuntu:latest

# INFO: Download mandatory software.
RUN apt update && apt install -y adduser

# INFO: Set workdir.
WORKDIR /home/ubuntu

# INFO: Download mandatory software.
RUN apt update && apt install -y curl

# INFO: Download the latest runner package.
RUN curl -o actions-runner-linux-x64-2.331.0.tar.gz -L https://github.com/actions/runner/releases/download/v2.331.0/actions-runner-linux-x64-2.331.0.tar.gz

# INFO: Validate the hash.
RUN echo "dccad95cd44acedf25f89ab03ec47d848fef27c9e8d445a194e3166a07b531b053e8efea7112644d03e0360f3bd108a9fbc596567290a1af773e60e6f8448f21  actions-runner-linux-x64-2.331.0.tar.gz" | sha512sum -c

# INFO: Extract the installer.
RUN mkdir -p actions-runner
RUN tar xzf actions-runner-linux-x64-2.331.0.tar.gz -C actions-runner
RUN rm -rf actions-runner-linux-x64-2.331.0.tar.gz

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