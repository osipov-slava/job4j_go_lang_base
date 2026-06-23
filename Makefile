# Переменные
GO := go
GO_PKG := ./cmd/app
BINARY_NAME := myapp
BUILD_DIR := ./bin

# Цель по умолчанию
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  test        - Run all tests"
	@echo "  coverage    - Run tests and generate HTML coverage report"
	@echo "  cover       - Alias for coverage"
	@echo "  lint        - Run golangci-lint"
	@echo "  build       - Build the application"
	@echo "  all         - Run lint, tests and coverage"
	@echo "  help        - Show this help"

# Сборка приложения
.PHONY: build
build:
	@echo "Building application..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(GO_PKG)
	@echo "✅ Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Запуск всех тестов
.PHONY: test
test:
	$(GO) test -v $(GO_PKG)

# Генерация отчёта о покрытии в формате HTML
.PHONY: coverage cover
coverage cover:
	$(GO) test -coverprofile=coverage.out $(GO_PKG)
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: file://$(shell pwd)/coverage.html"

# Вывод покрытия в терминал (опционально)
.PHONY: cover-report
cover-report:
	$(GO) test -cover $(GO_PKG)

# Проверка кода с помощью golangci-lint
.PHONY: lint
lint:
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "❌ golangci-lint is not installed. Please install it:"; \
		echo "   https://golangci-lint.run/usage/install/"; \
		exit 1; \
	fi
	golangci-lint run

# Запуск всех проверок
.PHONY: all
all: lint test coverage