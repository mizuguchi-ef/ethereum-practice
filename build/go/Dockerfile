FROM golang:1.23.1-alpine

# アップデートとgitのインストール
RUN apk update && apk add git

# ワーキングディレクトリの設定
WORKDIR /src

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /src