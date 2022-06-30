# goのベースイメージで1.18を使います。
FROM golang:1.18.0-alpine3.14

# ROOTという名前で、`/go/src/app`を定義します。
ENV ROOT=/go/src/app
# WORKDIRを/go/src/appにします。
WORKDIR ${ROOT}

#  gitをインストールします。（ライブラリインストール用）
RUN apk update && apk add git

# main.goを/go/src/appにコピーします。
# COPY ./main.go ${ROOT}
# go.mod, go.sumを/go/src/appにコピーします。
COPY go.mod go.sum ${ROOT}

# go mod tidyを実行します。
RUN go mod tidy

RUN go install github.com/cosmtrek/air@v1.27.3
CMD ["air", "-c", ".air.toml"]
