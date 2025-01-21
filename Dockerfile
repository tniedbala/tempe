FROM golang:1.23-bookworm

ARG ANTLR_VERSION
ENV ANTLR_VERSION=${ANTLR_VERSION:-4.13.2}
ENV ANTLR_JAR_PATH=/usr/local/bin/antlr-$ANTLR_VERSION.jar
ENV WORKSPACE=/workspace
ENV CLASSPATH=$ANTLR_JAR_PATH:$WORKSPACE/bin/grammar

WORKDIR $WORKSPACE
COPY . .
RUN apt update -y && apt install -y curl git default-jdk && \
    curl -L -o $ANTLR_JAR_PATH "https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar" && \
    echo 'alias antlr="java -jar $ANTLR_JAR_PATH"' >> ~/.bashrc && \
    echo 'alias antlr-test="java org.antlr.v4.gui.TestRig"' >> ~/.bashrc

ENTRYPOINT ["/bin/bash","scripts/docker-entrypoint.sh"]