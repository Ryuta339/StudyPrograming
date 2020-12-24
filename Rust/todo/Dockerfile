# Docker Official の Rust イメージを使う。
FROM rust:1.48 AS builder

# /todo でビルドを行う。
WORKDIR /todo

# Cargo.toml のみ先にイメージに追加。
COPY Cargo.toml Cargo.toml

# ビルドするために何もしないソースコードを入れておく。
RUN mkdir src
RUN echo "fn main() {}" > src/main.rs

# 上記で作成したコードと依存クレートをビルド。
RUN cargo build --release

# アプリケーションのコードをイメージにコピー。
COPY ./src ./src
COPY ./templates ./templates

# 先にビルドしたもののうち、アプリケーションにもののみを削除。
RUN rm -rf target/release/deps/todo*

# build
RUN cargo build --release

# 新しくリリース用のイメージを用意。
FROM debian:10.4 

# builder イメージから todo のみをコピーして /usr/local/bin に配置
COPY --from=builder /todo/target/release/todo /usr/local/bin/todo

# コンテナ起動時にWebアプリを実行。
CMD ["todo"]
