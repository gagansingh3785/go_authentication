SHELL = /bin/sh


start-local-server:
	source "./local.env" && go run .

start-prod-server:
	source "./../scripts/production.env" && go run .