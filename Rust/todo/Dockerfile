# Docker Official の Rust イメージを使う。
FROM rust:1.48

# /todo でビルドを行う。
WORKDIR /todo

# ビルドに必要なファイルをイメージにコピー。
COPY Cargo.toml Cargo.toml
COPY ./src ./src
COPY ./templates ./templates

# build
RUN cargo build --release

# パスの通った場所にインストール。
RUN cargo install --path .

# コンテナ起動時にWebアプリを実行。
CMD ["todo"]
