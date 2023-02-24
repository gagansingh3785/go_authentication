SHELL = /bin/sh

current_dir = $(shell pwd)


start-local-server:
	source "./local.env" && go run .

start-prod-server:
	source "$(current_dir)/../scripts/production.env" && go run .