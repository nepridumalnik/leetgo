.PHONY: easy medium hard

%:
	@:

easy:
	@ARG="$(filter-out $@,$(MAKECMDGOALS))"; \
	if [ -z "$$ARG" ]; then echo "Ошибка: Укажите URL. Пример: make easy https://leetcode.com"; exit 1; fi; \
	go run cmd/create_task/main.go -level easy -url "$$ARG"; \
	TASK_NAME=$$(echo "$$ARG" | sed -E 's|.*/problems/([^/]+).*|\1|' | tr '-' '_'); \
	echo "easy/$$TASK_NAME/solution.go"; \
	echo "easy/$$TASK_NAME/solution_test.go"

medium:
	@ARG="$(filter-out $@,$(MAKECMDGOALS))"; \
	if [ -z "$$ARG" ]; then echo "Ошибка: Укажите URL. Пример: make medium https://leetcode.com"; exit 1; fi; \
	go run cmd/create_task/main.go -level medium -url "$$ARG"; \
	TASK_NAME=$$(echo "$$ARG" | sed -E 's|.*/problems/([^/]+).*|\1|' | tr '-' '_'); \
	echo "medium/$$TASK_NAME/solution.go"; \
	echo "medium/$$TASK_NAME/solution_test.go"

hard:
	@ARG="$(filter-out $@,$(MAKECMDGOALS))"; \
	if [ -z "$$ARG" ]; then echo "Ошибка: Укажите URL. Пример: make hard https://leetcode.com"; exit 1; fi; \
	go run cmd/create_task/main.go -level hard -url "$$ARG"; \
	TASK_NAME=$$(echo "$$ARG" | sed -E 's|.*/problems/([^/]+).*|\1|' | tr '-' '_'); \
	echo "hard/$$TASK_NAME/solution.go"; \
	echo "hard/$$TASK_NAME/solution_test.go"
