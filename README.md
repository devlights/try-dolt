# try-dolt

This is my TUTORIAL project for dolt.

[dolt](https://github.com/dolthub/dolt) は、```Git for Data``` と書かれているように ```データベース + Git``` というイメージのソフトウェア。

通常のデータベースの様にクエリを発行してデータを操作することも可能で、データベースの変更情報をgitのようにバージョン管理することが出来る。

[dolt](https://github.com/dolthub/dolt) は、MYSQL互換のデータベースサーバーとしても動作可能。

## REFERENCES

- dolthub/dolt
  - https://github.com/dolthub/dolt
- dolt documents
  - https://docs.dolthub.com/
- Dolt で Ubie のマスターデータを管理してみる
  - https://zenn.dev/ubie_dev/articles/a0b74af675d4bc
- RDBのデータの構成管理ができるDoltとは
  - https://qiita.com/kitaj/items/4a72bdcc4c91aa31a81e
- Git for Data「Dolt」というDBの話
  - https://zenn.dev/masamiki/articles/e75ae440525ac0
- Git for Data、バージョン管理データベース Dolt が PostgreSQL 仕様になる
  - https://www.infoq.com/jp/news/2023/12/DoltgreSQL-git-for-data-postgres/
- Announcing DoltgreSQL
  - https://www.dolthub.com/blog/2023-11-01-announcing-doltgresql/
- dolthub/doltgresql (*1)
  - https://github.com/dolthub/doltgresql

---

(*1) perplexityで要約した内容

> DoltHubは、DoltgreSQLの開発を開始したことを発表しました。DoltはMySQLの構文を使用するバージョン管理データベースで、Gitのようなブランチやマージ、差分などの機能をデータベースに適用しています。DoltgreSQLは、PostgreSQLの構文をサポートする新しい製品で、コマンドライン機能はなく、データベースサーバーとしての使いやすさに重点を置いています。

> DoltgreSQLは、PostgreSQLサーバーをエミュレートし、受信したコマンドをDoltサーバーに渡すことで動作します。このアプローチにより、開発が迅速に進められ、Doltの安定した機能を活用できます。

> DoltgreSQLの使用は非常に簡単で、最新リリースをダウンロードし、バイナリをPATHに追加し、データベースを作成して接続するだけです。ただし、現在は非常に初期段階であり、多くのエラーや問題が予想されます。頻繁にリリースが行われる予定で、最新の進捗を確認するにはソースからビルドするのが最適です。

> 今後の開発に向けて、ユーザーからのフィードバックや貢献を歓迎しています。TwitterやDiscordで意見を共有することができます。
