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

- Set Environmental valuable for Dev

```
touch build/.env.local
vi build/.env.local

# Copy contents from ()
# and save with command below
:wq
```

- Start Docker

```
docker-compose build
docker-compose up -d
```
