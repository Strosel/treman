FROM golang

RUN go get -d github.com/strosel/treman
RUN go get gioui.org/cmd/gogio

RUN apt-get update
RUN apt-get install --yes \
    libxkbcommon-x11-dev \
    libgles2-mesa-dev \
    curl \
    jq

ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]