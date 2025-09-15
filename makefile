build:
	docker-compose up --build

run:
	docker-compose up

config:
	git config --global user.name alseitkad
	git config --global user.email alisherseitkadyr@gmail.com

push:
	@if [ -z "$(m)" ]; then \
		echo "❌ Please provide a commit message: make push m='Your message here'"; \
	else \
		git add . && git commit -m "$(m)" && git push; \
	fi

clean:
	docker-compose down -v

updatego:
	chmod +x s.sh
	./s.sh
