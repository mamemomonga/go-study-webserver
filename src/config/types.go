package config

/*
type でYAMLとデータ構造をあわせる
キーは先頭を大文字の項目に読み込まれる

https://godoc.org/gopkg.in/yaml.v2#Unmarshal
https://ota42y.com/blog/2014/12/03/go-yaml-struct/
http://tweeeety.hateblo.jp/entry/2017/06/04/231043

*/

type WebServer struct {
	Listen string
}

type Config struct {
	Webserver WebServer
}

