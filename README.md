# Go入門: ウェブサーバをつくる

* シンプルなウェブサーバ
* YAML形式設定ファイル
* オブジェクト指向風の構造
* マルチアーキテクチャービルド

# 実行環境

* macOS High Sierra
* go version go1.10.1 darwin/amd64

# 準備

	$ cd src
	$ go get -v
	$ cd ..

# 実行

	$ go run src/main.go etc/config.yaml

http://localhost:8000/ でアクセスできる

# マルチアーキテクチャービルド

	$ ./build.sh

macOSでビルド済みバイナリの実行例

	$ build/webserver.darwin-amd64 etc/config.yaml

