start-app: start-api \
 start-ui

start-api:
	go run .

start-ui:
	cd ui && npm start

