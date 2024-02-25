run:
	go run ./cmd/mithril/mithril.go

api-new-user:
	curl --request POST -sL \
	     --url 'http://localhost:8081/api/users'\
	     --data '{"username": "foo", "password": "bar"}'

api-get-root:
	curl --request GET -sL \
	     --url 'http://localhost:8081/'
