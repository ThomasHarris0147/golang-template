ENV ?= local

.PHONY: makemigrations applymigrations migrate-status migrate-lint schema-diff run

run:
	air serve

help:
	@echo "Atlas Migration Commands:"
	@echo "  make makemigrations NAME=<name>  Generate a new migration"
	@echo "  make applymigrations             Apply pending migrations"
	@echo "  make migrate-status              Show migration status"
	@echo "  make migrate-lint                Lint migrations"
	@echo "  make schema-diff                 Show schema differences"
	@echo "  make migrate-dry-run             Preview migration changes"
	@echo ""
	@echo "Examples:"
	@echo "  make makemigrations NAME=initial"
	@echo "  make makemigrations NAME=add_username"

makemigrations:
	@if [ -z "$(NAME)" ]; then \
		echo "❌ Error: Migration name required"; \
		echo "Usage: make makemigrations NAME=<migration_name>"; \
		echo "Example: make makemigrations NAME=initial"; \
		exit 1; \
	fi
	atlas migrate diff $(NAME) --env $(ENV)

applymigrations:
	atlas migrate apply --env $(ENV)

migrate-status:
	atlas migrate status --env $(ENV)

migrate-lint:
	atlas migrate lint --env $(ENV)

schema-diff:
	atlas schema diff --env $(ENV)

migrate-dry-run:
	atlas migrate apply --env $(ENV) --dry-run

verify-models:
	go run ariga.io/atlas-provider-gorm load --path ./models --dialect sqlite