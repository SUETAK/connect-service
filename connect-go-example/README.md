# proto の更新
/connect-go-example
```sh
 buf generate
```

# ORM の導入モチベーション
- DBへの接続はORM の[Bunライブラリ](https://bun.uptrace.dev/)を使用する
- ORM SQLファーストなので、ORM独自の仕様をあまり考えずに記載できることを優先する
