# Golang API sample

## 機能概要

GolangでRESTサーバーを起動し、POSTメソッドを受け付けるAPI。

1. httpieで`localhost:9999/hello`にアクセス
2. JSON形式で`Name=xxxx`をリクエスト
3. `Hello, xxxx`でレスポンスを受け取る

## 環境

 - go1.9.3
 - httpie 1.0.2

## 使用パッケージ

 - log
 - net/http
 - github.com/ant0ine/go-json-rest/rest