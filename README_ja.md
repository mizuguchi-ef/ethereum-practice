# ethereum-practice

- [English](README.md)

ethereum-practice の開発環境

## ローカル環境構築手順

## 前提条件

- Docker インストール
  https://docs.docker.com/docker-for-mac/install
- docker-compose インストール
  https://docs.docker.jp/compose/install.html
- Git インストール
  `brew install git`

## 構築手順

- ethereum-practice をクローンして該当のディレクトリへ移動

```
git clone -b main https://github.com/mizuguchi-ef/ethereum-practice.git
cd ethereum-practice
```

- Docker 起動

```
docker-compose build
docker-compose up -d
```

## コンテナに入り、geth を起動

```
docker exec -it geth /bin/sh
geth --http --http.addr "0.0.0.0" --http.port "8545" --http.api "eth,net,web3" --ws --ws.addr "0.0.0.0" --ws.port "8546" --ws.api "eth,net,web3" --syncmode "snap"
```

(http://localhost:8545)にアクセスし、RPC が起動しているか確認
