【概要】
ブログサービスの作成

【アーキテクチャ】
 - バックエンド     ：Golang
 - フロントエンド   ：Next.js, Typescript
 - モデル           ：MySQL
 - 基盤             ：Docker（バックエンド、フロントエンド、モデル層それぞれをコンテナ管理）

詳細
-----------------------------------------------------------------------------
【バックエンド】
- 概要
    ブログサービス管理用API
    記事の投稿、投稿されたデータの取得などの機能を搭載する

- パッケージ構成
    api           ：ルーティング処理を担当
    controllers   ：リクエストに対するハンドラの管理を担当
    services      ：ビジネスロジックを担当
    repositories  ：DBアクセス処理を担当
    models        ：各テーブルの構造体定義を管理を担当
    errors        ：独自エラー型の管理を担当

- 学習技術（抜粋）
    - 基礎文法(各種変数の扱い、条件分岐、ループ処理)
    - gorilla/muxパッケージによるルーティング
    - mysqlドライバーによるDBアクセス処理
    - インタフェースによる抽象化
    - 独自エラー型によるエラー処理

- 参考
　https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx?productVariantID=dvjtgpjw8VDTXNqKaanTVi
-----------------------------------------------------------------------------
【フロントエンド】※予定（構想のみ、未作成）
    ブログサービス画面表示
    記事の投稿、投稿されたデータの参照などの機能を搭載する
