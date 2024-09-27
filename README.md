# ethereum-practice

- [Japanese](README_ja.md)

For ethereum-practice development environment

## Procedure for building your local environment

## Prerequisite

- Install Docker
  https://docs.docker.com/docker-for-mac/install
- Install docker-compose
  https://docs.docker.jp/compose/install.html
- Install Git
  `brew install git`

## Construction procedure

- Clone ethereum-practice and go to that directory

```
git clone -b main https://github.com/mizuguchi-ef/ethereum-practice.git
cd ethereum-practice
```

- Start Docker

```
docker-compose build
docker-compose up -d
```

## Enter the container and Start Geth

```
docker exec -it geth /bin/sh
geth --http --http.addr "0.0.0.0" --http.port "8545" --http.api "eth,net,web3" --ws --ws.addr "0.0.0.0" --ws.port "8546" --ws.api "eth,net,web3" --syncmode "snap"
```

Access to (http://localhost:8545) to check RPC is activated
