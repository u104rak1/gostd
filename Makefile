.PHONY: up run help

up: ## コンテナを起動
	docker-compose up -d

rebuild: ## コンテナを再構築
	docker-compose up -d --build

down: ## コンテナを停止
	docker-compose down

run: ## アプリケーションを実行
	go run cmd/gostd/main.go

help: ## ヘルプを表示
	@grep -E '(^##|^[a-zA-Z_-]+:.*?##)' $(MAKEFILE_LIST) | \
		awk '/^##/ {print substr($$0, 4)} /^[a-zA-Z_-]+:/ {split($$0, a, ":.*?## "); printf "\033[36m%-30s\033[0m %s\n", a[1], a[2]}'