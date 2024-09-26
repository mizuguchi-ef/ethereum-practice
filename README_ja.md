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

- 環境変数を設定

```
touch build/.env.local
vi build/client/.env.local

# () から内容をコピーして以下コマンドで保存
:wq
```

- Docker 起動

```
docker-compose build
docker-compose up -d
```
