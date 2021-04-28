# microcms-sdk-generator

## Usage

1. microCMS のダッシュボードから[スキーマをエクポート](https://document.microcms.io/manual/export-and-import-api-schema)してください。

2. 下記の例とリファレンスを参照しながら `metadata.yml` を記述してください。

<details>
  <summary>例</summary>

```yaml
info:
  service-id: my-blog-service
  base-path: ./schemas/ # カレントディレクトリを指定する場合、省略可能

apis:
  # 最も単純な例
  - name: users

  # 最も複雑な例
  - name: articles
    path: ./articles.json
    relations:
      author: user
    custom-fields:
      ogp: meta

  # オブジェクト形式の API の例
  - name: popular-articles
    form: object
```
</details>

<details>
  <summary>リファレンス</summary>

### `info.service-id`

サービスの ID を指定してください。

API のエンドポイントが `https://my-service-name.microcms.io/api/v1/my-api-name` である場合は、 `my-service-name` を指定してください。

### `info.base-path`

microCMS のスキーマファイルが配置されているディレクトリを指定してください。

このフィールドは省略可能です。デフォルト値は `.` です。

### `apis.*.name`

API のエンドポイント名を指定してください。

API のエンドポイントが `https://my-service-name.microcms.io/apis/sellers/my-api-name` である場合は、 `my-api-name` を指定してください。

### `apis.*.path`

microCMS のスキーマファイルが配置されているファイルのパスを指定してください。

このフィールドは省略可能です。デフォルト値は `${info.base-path}/${apis.*.name}.json` のような値になります。

### `apis.*.form`

API 作成時に指定したレスポンスの形式を指定してください。

このフィールドには `list`, `object` のどちらかを指定してください。

このフィールドは省略可能です。デフォルト値は `list` です。

### `apis.*.relations`

API のリレーションを連想配列で指定してください。
連想配列のキーには API のエンドポイント名 (つまり `apis.*.name`) を指定してください。
連想配列の値にはフィールド ID を指定してください。

例えば `articles` というオブジェクトが `users` というオブジェクトを `author` というフィールド名で持っている場合、以下のように記述してください。

```yaml
apis:
  users:
    # other fields...
  articles:
    relations:
      author: user
      # other fields...
```

他の API とのリレーションが無い場合、このフィールドは省略可能です。

カスタムフィールドは `apis.*.relations` で指定する必要はありません。

### `apis.*.custom-fields`

API で使用しているカスタムフィールドの使用箇所を指定する連想配列です。
連想配列のキーにはフィールド ID を指定してください。
連想配列の値にはカスタムフィールド ID を指定してください。

例えば `users` というオブジェクトが `sns` というカスタムフィールドを `twitter`, `facebook` というフィールド名で持っている場合、以下のように記述してください。

```yaml
apis:
  contents:
    custom-fields:
      twitter: sns
      facebook: sns
    # other fields...
```

カスタムフィールドが存在しない場合、このフィールドは省略可能です。
</details>

3. `docker run lambdasawa/microcms-sdk-generator` を実行してください。 

<details>
  <summary>オプションの指定例</summary>

```sh
# このようなディレクトリ構造があるとき…
$ tree
  .
  └── microcms
      ├── metadata.yml         # ステップ 2 で作成したファイル 
      └── schemas              # microCMS からエクスポートしたスキーマファイルがあるディレクトリ 
          ├── articles.json
          ├── setting.json
          └── users.json

# このようなオプションでコマンドを実行すると…
$ docker run \
    --rm \
    --volume $PWD/microcms/:/app/microcms/ \
    --env METADATA_PATH=/app/microcms/metadata.yml \
    --env OUTPUT_PATH=/app/microcms/ \
    lambdasawa/microcms-sdk-generator \
    --generator-name typescript-fetch \
    --additional-properties=typescriptThreePlus=true,allowUnicodeIdentifiers=true

# このようにファイルが生成されます。
$ tree
  .
  └── microcms
      ├── apis
      │   ├── ArticlesApi.ts
      │   ├── SettingApi.ts
      │   ├── UsersApi.ts
      │   └── index.ts
      ├── index.ts
      ├── metadata.yml
      ├── models
      │   ├── Articles.ts
      │   ├── ArticlesCreateRequest.ts
      │   ├── ArticlesList.ts
      │   ├── ArticlesPatchRequest.ts
      │   ├── CommonUpdateResult.ts
      │   ├── Setting.ts
      │   ├── SettingCreateRequest.ts
      │   ├── SettingList.ts
      │   ├── SettingPatchRequest.ts
      │   ├── Users.ts
      │   ├── UsersCreateRequest.ts
      │   ├── UsersIcon.ts
      │   ├── UsersList.ts
      │   ├── UsersPatchRequest.ts
      │   ├── UsersSns.ts
      │   └── index.ts
      ├── runtime.ts
      └── schemas
          ├── articles.json
          ├── setting.json
          └── users.json
```

末尾の引数は [`openapi-generator` の `generate` コマンド](https://openapi-generator.tech/docs/usage#generate) に渡されます。
</details>

## TODO

- [ ] `examples/` 追加
- [ ] ユニットテスト追加
- [ ] `orders`, `fields` の値としてネストしたフィールドのサポート
- [ ] english document
