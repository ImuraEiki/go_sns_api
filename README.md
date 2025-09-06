## go_sns_api
React学習のために開発中の[react-sns](https://github.com/ImuraEiki/react_sns)のAPIです。

## 準備

### インストールする
sudo apt install golang-go

### Dockerイメージをビルド
docker build -f Dockerfile.prod -t my-gin-app .

### コンテナを起動（ローカルの8080をコンテナの8080にマッピング）
docker run -p 8080:8080 my-gin-app

### 開発用コンテナの場合
docker-compose up --build
http://localhost:8080/api/users

#### Tips
- `no required module provides package github.com/gin-contrib/cors`と出たらローカルで`go get github.com/gin-contrib/cors`を実行
- `error obtaining VCS status`と出たら`git config --global --add safe.directory /home/eiki/workspace/go_api(作業ディレクトリ)`