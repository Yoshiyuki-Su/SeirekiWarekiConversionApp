# 西暦、和暦変換アプリ

## 仕様

西暦を入れた場合は和暦を表示。  
和暦を入れた場合は西暦へ変換して表示する。

ex:
go run main.go -s 20220729
令和4年7月29日

go run main.go -w 令和4年7月29日
20220729

## 使い方

フラグ

```
-s <西暦>
result: 和暦
-w <和暦>
result: 西暦
```

## Init

```
$ go mod init SeirekiWarekiConversionApp
go: creating new go.mod: module SeirekiWarekiConversionApp
```

## 実施方法

```
$ go run main.go -seireki 20210701
令和3年

$ go run main.go -s 20210701      
令和3年
```

## 参考

> https://qiita.com/nrnrk/items/0ad21cfa8c4c8357dae1
> 