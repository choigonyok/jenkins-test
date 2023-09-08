FROM golang:latest

WORKDIR /app

COPY .go mod ./

RUN go mod download

COPY ./ ./

RUN go build main.go

EXPOSE 8080

CMD [ "./main" ]

# image build시 docker build -t TAG -f build/Dockerfile . 로 빌드해야 부모요소도 컨테이너로 복사가 가능하다.