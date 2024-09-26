FROM golang-1.23.1-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# ワーキングディレクトリの設定
WORKDIR /go/src/github.com/ethereum-practice
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com/ethereum-practice
# パッケージのインポート
RUN go install golang.org/x/tools/cmd/goimports@latest