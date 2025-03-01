# ドメインモデル: Task

## **エンティティ**

### **Task**

タスクを表すエンティティ

| フィールド  | 型                 | 説明                                              |
| ----------- | ------------------ | ------------------------------------------------- |
| id          | UUIDv7 (string)    | タスクを一意に識別する ID（時系列順にソート可能） |
| title       | Title              | タスクのタイトル（最大 20 文字）                  |
| description | Description (任意) | タスクの詳細説明（最大 255 文字）                 |
| priority    | Priority           | タスクの優先度 (`low medium high`)                |
| stage       | Stage              | タスクの進行状況 (`TODO → IN_PROGRESS → DONE`)    |
| updated_at  | datetime           | タスクの最終更新日時                              |

---

### **Priority**

タスクの優先順位を定義

| 値     | 説明     |
| ------ | -------- |
| low    | 低優先度 |
| medium | 中優先度 |
| high   | 高優先度 |

---

### **Stage**

タスクの進行状況

| 値          | 説明                 |
| ----------- | -------------------- |
| TODO        | 未着手（デフォルト） |
| IN_PROGRESS | 進行中               |
| DONE        | 完了                 |

---

## **制約事項**

- `title` は **必須**
- `priority` のデフォルト値は `medium`
- `stage` は `TODO` → `IN_PROGRESS` → `DONE` の順番でしか変更できない

---

### **Tasks**

タスクをソート、フィルタリングする役割を持つ
|フィールド|型|説明|
| tasks |Task[]| タスク一覧 |
