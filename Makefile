.PHONY: wire

# wire 自动依赖注入生成
wire:
	@echo ">>> Generating wire injectors..."
	cd ./internal/di && wire
	@echo ">>> Wire generation complete."