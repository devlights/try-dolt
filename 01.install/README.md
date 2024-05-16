# インストールについて

[dolt](https://github.com/dolthub/dolt)のインストールについては、以下に記載があります。

- https://github.com/dolthub/dolt?tab=readme-ov-file#installation

[dolt](https://github.com/dolthub/dolt)自身はGoで作成されているので、シングルバイナリとなっています。

なので、一つのバイナリをパスを通ったところに置けば終わりです。

## Windows の場合

[Scoop](https://scoop.sh/) でインストールするのが楽です。

```sh
$ scoop install dolt
```

## Macの場合

Homebrewでインストールするのが楽です。

```sh
$ brew install dolt
```

## Linuxの場合

以下のコマンドを実行すると ```/usr/local/bin``` の下に配置されます。

```sh
$ sudo bash -c 'curl -L https://github.com/dolthub/dolt/releases/latest/download/install.sh | bash'
```

# バージョンの確認

```sh
$ dolt version
dolt version 1.38.0
```

のように表示されればインストールは完了です。
