test:
	go run main.go -data.path=example/data.yml -template.path=example/config.yml.tpl -output.path=example/config.yml
	cat example/config.yml
