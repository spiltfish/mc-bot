FROM golang
RUN mkdir /mc-bot
ADD . /mc-bot
WORKDIR /mc-bot
RUN go get github.com/bwmarrin/discordgo && github.com/spiltfish/mc-worker-sdk && go get gopkg.in/yaml.v2
RUN go build -o main .
CMD ["/mc-bot/main"]
