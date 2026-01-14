FROM ubuntu:25.04

# INFO: Install base packages.
RUN apt-get update && apt-get install -y bash-completion

# INFO: Install development packages.
RUN apt-get update && apt-get install -y git vim