.PHONY: build up down run stop rerun mod-tidy help

build: ## Docker イメージをビルド
	docker-compose build

up: ## Docker コンテナを起動
	docker-compose up -d

down: ## Docker コンテナを停止・削除
	docker-compose down

run: ## DelveデバッガーとAPIサーバーを同時起動
	docker-compose exec -d backend dlv debug ./cmd/gostd --build-flags=-buildvcs=false --headless --listen=:2345 --api-version=2 --accept-multiclient --continue
	sleep 2

stop: ## DelveとAPIサーバーを完全停止（コンテナは永続化維持）
	docker-compose exec backend pkill -f "go run" || true
	docker-compose exec backend pkill -f "__debug_bin" || true
	docker-compose exec backend pkill -f "gostd" || true
	docker-compose exec backend pkill -f "dlv" || true
	sleep 2

rerun: ## make stop → make run
	make stop && make run

mod-tidy: ## Go modulesを更新
	docker-compose exec backend go mod tidy

help: ## ヘルプを表示
	@grep -E '(^##|^[a-zA-Z_-]+:.*?##)' $(MAKEFILE_LIST) | \
		awk '/^##/ {print substr($$0, 4)} /^[a-zA-Z_-]+:/ {split($$0, a, ":.*?## "); printf "\033[36m%-16s\033[0m %s\n", a[1], a[2]}'


