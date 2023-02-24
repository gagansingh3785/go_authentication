SHELL = /bin/bash

current_dir = $(shell pwd)


start-local-server:
	source "./local.env" && go run .

start-prod-server:
	source "$(current_dir)/../scripts/production.env" && $(current_dir)/go_authentication