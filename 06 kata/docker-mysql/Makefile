build:
	@docker build -t go-mysql .

start:
	@docker run -p 3306:3306 --name go-mysql go-mysql --default-authentication-plugin=mysql_native_password

stop:
	@docker stop go-mysql

clean: stop
	@docker rm go-mysql