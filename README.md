# CFnドリフト警察
## 実装の前提
ドメイン駆動開発チックに作成しています（この規模のアプリケーションであれば、この作りにする必要はありませんが、学習のため）
また、試しにGoで書いてみた、というものなので、Goのよくある書き方とは異なる書き方になっています。

## 仕様
CloudFormationの各スタックに対して、ドリフトの検出を行い、ドリフトが発生している場合は、通知を行うシステムです。

# 開発
## 前提条件
|#|ソフトウェア|バージョン|
| ---- | ---- | ---- |
|1|Go|1.x系|

## テスト
ユニットテストはモックをインジェクションすることで実現する自動テストが可能です。クライアント層はモックされているため、aws権限の準備なくテスト実行できます。

```
$ make init
$ make test
```

### モックの追加方法
以下のコマンドを使用して、インタフェースからモックを生成することができます。

```
$ mockgen -source 生成元となるインタフェースのパス -destination モックを生成する対象となるパス
```

## ローカルデバッグ
特別な設定を行わない限りは、クライアント層がモックされるため、aws権限の準備なく実行できます。
以下のコマンドで行うことができます。

```
$ make init
$ make invokeCheck
$ make invokeAlert
```

このリポジトリをチェックアウトしたら、パッケージのインストールとローカル実行を試してみてください。
その他`make`でできることは`Makefile`をご参考ください。