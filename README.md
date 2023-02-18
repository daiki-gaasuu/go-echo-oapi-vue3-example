# Go＋Vue3ローカル検証用

## バックエンド

- Go1.17.6

## フロントエンド

- Vite4.1.0
- Vuetify3
- Vue3
- Node19.6.119.6.1
- yarn4.0.0-rc.39

## OpenAPI

- バックエンド
  - https://github.com/deepmap/oapi-codegen
  - プロジェクトルートにて`make oapi-generate-backend`
- フロントエンド
  - プロジェクトルートにて`make openapi-generate-frontend`
  - https://github.com/OpenAPITools/openapi-generator
  
## 起動方法
1. 上記のOpenAPIの生成
2. ```cd server```
3. ```go run main.go```
4. clientルートにて```yarn```
5. yarn run dev
