# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/go/.devcontainer/base.Dockerfile
ARG VARIANT="1.19-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; \
  then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# Get JDK11 for ANTLR4.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
  && apt-get -y install --no-install-recommends openjdk-11-jdk wget

# Setup the latest verison of ANTLR4.
RUN wget https://www.antlr.org/download/antlr-4.11.1-complete.jar -O /usr/local/lib/antlr.jar
RUN echo "#/bin/sh" > /usr/local/bin/antlr && \
  echo "java -Xmx500M -cp /usr/local/lib/antlr.jar org.antlr.v4.Tool \$@" >> /usr/local/bin/antlr && \
  chmod +x /usr/local/bin/antlr

# Install Go tooling for VS Code as vscode user.
USER vscode
RUN go install github.com/josharian/impl@latest && \
  go install github.com/cweill/gotests/gotests@latest && \
  go install github.com/fatih/gomodifytags@latest && \
  go install github.com/haya14busa/goplay/cmd/goplay@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest && \
  go install honnef.co/go/tools/cmd/staticcheck@latest && \
  go install golang.org/x/tools/gopls@latest
