FROM golang:1.23.1-alpine

# アップデートとgitのインストール
RUN apk update && apk add git

# ワーキングディレクトリの設定
WORKDIR /go/src/github.com/practice

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com/practice