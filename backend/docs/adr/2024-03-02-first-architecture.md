# アーキテクチャ設計

## コンセプト

スケラーブルな TODO リストを作るために、過剰ではあるがヘキサゴナルアーキテクチャを使用する

### 各層の役割

#### 外部層

**Web アダプター**

- API のエンドポイントを提供し、リクエストをユースケースに橋渡しする役割。
- ポートを使ってユースケースを呼び出す。（直接ビジネスロジックにアクセスしない）
- 今回は認証を含まないので、サニタイズなどを行う

**リポジトリ**

- DB の CRUD 処理を担当する。
- DB の具体的な実装（SQL, NoSQL, etc.）をカプセル化し、ユースケースはインターフェースを通じてアクセスする。 -「リスコフの置換原則」を守るために、リポジトリのインターフェースを定義する。

#### アプリケーション層

**ドメイン**

- ビジネスルールやドメインの不変条件を管理する。
- 他の層には依存しない（特に外部層への依存は厳禁）。
- ドメインオブジェクトが正しい状態を保持するように、コンストラクタやメソッドでバリデーションを行う。

**ユースケース**

- ビジネスロジックの流れを定義する
- トランザクションもここで定義する
- 複数のリポジトリやサービスを組み合わせる処理をここで記述する。

**ポート層**

- ユースケースごとの入り口である。
- 入力モデルもここで定義するものとする
- 入力モデル（DTO）もここで定義する。
