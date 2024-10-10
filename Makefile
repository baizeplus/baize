

init:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
wire:
	cd app/ && wire

swag:
	cd app/ && swag  init