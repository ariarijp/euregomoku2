euregomoku2
=====================

## 取り組んだこと

[Network programming with Go](https://jan.newmarch.name/go/)の[HTTP](https://jan.newmarch.name/go/http/chapter-http.html)を読みつつ写経というかコピペしつつコードを読んだ。Proxy handlingあたりは飛ばした。

掲載されてたサンプルを全部同じディレクトリに入れておきたかったので、commandパッケージを定義して、その中に各サンプルを入れるようにした。

## Usage

```
$ go run main.go head [url]
$ go run main.go get [url]
$ go run main.go client_get [url]
$ go run main.go file_server [path]
$ go run main.go print_env [path]
$ go run main.go server_handler
```

## License

MIT

## Author

[Takuya Arita](https://github.com/ariarijp)
