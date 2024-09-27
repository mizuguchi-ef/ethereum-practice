# ベースイメージ
FROM ethereum/client-go:latest

# 環境変数の設定
ENV DATA_DIR=/root/.ethereum GOROOT=go

# ポートを公開
EXPOSE 30303 8545 8546

# Geth の起動
CMD ["geth", "--http", "--http.addr", "0.0.0.0", "--http.port", "8545", "--http.api", "eth,net,web3", "--ws", "--ws.addr", "0.0.0.0", "--ws.port", "8546", "--ws.api", "eth,net,web3", "--syncmode", "snap"]