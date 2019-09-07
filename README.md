# auth-sample-jwt

これは、json web token(jwt)を利用した、ログインのバックエンド側のサンプルプロジェクトです。

# 動作環境

- windowsで開発していますが、goで作っているのでどの環境でも動作するはずです。

# バックエンドの起動

- 最初に一回だけ `make keys` でtoken署名用のkeyファイルを作成します。
- `src/main.go を実行します。`

# dockerで起動する

- `make deploy`で ${HOME}/auth-sample-jwt にdocker起動用の環境を構築します。
- `make start` または、${HOME}/auth-sample-jwt で `docker-compose up -d` でバックエンドが起動します。
