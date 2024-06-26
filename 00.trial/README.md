# doltのお試し

入門として、お試しでdoltで遊んでみた結果を以下にメモしておきます。

本ディレクトリのTaskfileを実行すると、以下のように流れます。

コマンドの詳細は [Taskfile.yml](./Taskfile.yml) を参照ください。

1. ```dolt```の初期化
1. データベース作成
1. テーブル作成
1. データ登録
1. ```dolt```のバージョン管理を使ってデータをステージング
1. ```dolt```のバージョン管理を使ってデータをコミット
1. ```dolt```のバージョン管理を使って変更履歴のログを確認
1. ```dolt```をMYSQL互換サーバとして起動
1. Goプログラムから接続して処理
1. Goプログラムで登録した変更分をバージョン管理

```sh
$ task
task: [delete] rm -rf ./.dolt
task: [delete] rm -rf ./.tmp
task: [delete] rm -rf ./db0
task: [delete] rm -f ./app

task: [initdb] dolt init
Successfully initialized dolt data repository.

task: [initdata] dolt sql --query "CREATE DATABASE db0;"
Query OK, 1 row affected (0.01 sec)
task: [initdata] dolt --use-db=db0 sql --query "CREATE TABLE tableA (id int primary key, name varchar(50));"
task: [initdata] dolt --use-db=db0 sql --query "INSERT INTO tableA (id, name) VALUES (1, 'gitpod1');"
Query OK, 1 row affected (0.00 sec)

task: [commit] dolt --use-db=db0 status
On branch main
Untracked tables:
  (use "dolt add <table>" to include in what will be committed)
        new table:        tableA
task: [commit] dolt --use-db=db0 add tableA
task: [commit] dolt --use-db=db0 diff --cached
diff --dolt a/tableA b/tableA
added table
+CREATE TABLE `tableA` (
+  `id` int NOT NULL,
+  `name` varchar(50),
+  PRIMARY KEY (`id`)
+) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_bin;
+---+----+---------+
|   | id | name    |
+---+----+---------+
| + | 1  | gitpod1 |
+---+----+---------+
task: [commit] dolt --use-db=db0 commit -m "Add tableA"
commit 73vbfbrkti8n6iaahjcr02sicp8iuk7a (HEAD -> main) 
Author: gitpod <gitpod@example.com>
Date:  Wed May 15 10:25:41 +0000 2024

        Add tableA

task: [commit] dolt --use-db=db0 log --oneline
73vbfbrkti8n6iaahjcr02sicp8iuk7a (HEAD -> main) Add tableA
qh5d91qj70fmphsbhubq6mu0hucsmocr Initialize data repository

task: [start-server] dolt sql-server &
task: [start-server] sleep 1
Starting server with Config HP="localhost:3306"|T="28800000"|R="false"|L="info"|S="/tmp/mysql.sock"
INFO[0000] Server ready. Accepting connections.         
WARN[0000] secure_file_priv is set to "", which is insecure. 
WARN[0000] Any user with GRANT FILE privileges will be able to read any file which the sql-server process can read. 
WARN[0000] Please consider restarting the server with secure_file_priv set to a safe (or non-existent) directory. 

task: [run-app] go build -o app
task: [run-app] for i in {1..3}; do ./app; done
INFO[0001] NewConnection                                 DisableClientMultiStatements=false connectionID=1
1:gitpod1
MAX(id)=1
INFO[0001] ConnectionClosed                              connectionID=1
INFO[0001] NewConnection                                 DisableClientMultiStatements=false connectionID=2
1:gitpod1
2:gitpod2
MAX(id)=2
INFO[0001] ConnectionClosed                              connectionID=2
INFO[0001] NewConnection                                 DisableClientMultiStatements=false connectionID=3
1:gitpod1
2:gitpod2
3:gitpod3
MAX(id)=3
INFO[0001] ConnectionClosed                              connectionID=3

task: [kill-server] pkill dolt
INFO[0001] Server closing listener. No longer accepting connections. 

task: [commit-changed] dolt --use-db=db0 status
On branch main

Changes not staged for commit:
  (use "dolt add <table>" to update what will be committed)
  (use "dolt checkout <table>" to discard changes in working directory)
        modified:         tableA
task: [commit-changed] dolt --use-db=db0 diff
diff --dolt a/tableA b/tableA
--- a/tableA
+++ b/tableA
+---+----+---------+
|   | id | name    |
+---+----+---------+
| + | 2  | gitpod2 |
| + | 3  | gitpod3 |
| + | 4  | gitpod4 |
+---+----+---------+
task: [commit-changed] dolt --use-db=db0 commit -am "Update"
commit d6pbvq5rt3phdrbllno0app67tntqdl8 (HEAD -> main) 
Author: gitpod <gitpod@example.com>
Date:  Wed May 15 10:25:42 +0000 2024

        Update

task: [commit-changed] dolt --use-db=db0 log --oneline
d6pbvq5rt3phdrbllno0app67tntqdl8 (HEAD -> main) Update
73vbfbrkti8n6iaahjcr02sicp8iuk7a Add tableA
qh5d91qj70fmphsbhubq6mu0hucsmocr Initialize data repository
task: [commit-changed] dolt --use-db=db0 diff HEAD~2 HEAD
diff --dolt a/tableA b/tableA
added table
+CREATE TABLE `tableA` (
+  `id` int NOT NULL,
+  `name` varchar(50),
+  PRIMARY KEY (`id`)
+) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_bin;
+---+----+---------+
|   | id | name    |
+---+----+---------+
| + | 1  | gitpod1 |
| + | 2  | gitpod2 |
| + | 3  | gitpod3 |
| + | 4  | gitpod4 |
+---+----+---------+
```