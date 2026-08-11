.PHONY: easy medium hard

# Заглушка, которая перехватывает URL (и любые другие аргументы после пробела)
# и не дает make упасть с ошибкой "target pattern contains no '%'" из-за двоеточия в https://
%:
	@:

easy:
	@ARG="$(filter-out $@,$(MAKECMDGOALS))"; \
	if [ -z "$$ARG" ]; then echo "Ошибка: Укажите URL. Пример: make easy https://leetcode.com"; exit 1; fi; \
	go run cmd/create_task/main.go -level easy -url "$$ARG"

medium:
	@ARG="$(filter-out $@,$(MAKECMDGOALS))"; \
	if [ -z "$$ARG" ]; then echo "Ошибка: Укажите URL. Пример: make medium https://leetcode.com"; exit 1; fi; \
	go run cmd/create_task/main.go -level medium -url "$$ARG"

hard:
	@ARG="$(filter-out $@,$(MAKECMDGOALS))"; \
	if [ -z "$$ARG" ]; then echo "Ошибка: Укажите URL. Пример: make hard https://leetcode.com"; exit 1; fi; \
	go run cmd/create_task/main.go -level hard -url "$$ARG"
