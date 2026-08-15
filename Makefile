.PHONY: easy medium hard

%:
	@:

easy:
	@go run cmd/create_task/main.go -level easy -url "$(filter-out $@,$(MAKECMDGOALS))"

medium:
	@go run cmd/create_task/main.go -level medium -url "$(filter-out $@,$(MAKECMDGOALS))"

hard:
	@go run cmd/create_task/main.go -level hard -url "$(filter-out $@,$(MAKECMDGOALS))"