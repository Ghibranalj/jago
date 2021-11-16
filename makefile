

build: container
	cd jago_bot && go build . && mv jago_bot ../jago
container :
	sudo docker build -t jago_compiler ./jago_compiler
run: build
	./run.sh &>.log