.PHONY: init
## init: init environment
init:
	pip install -r requirements.txt

.PHONY: format
## format: format files
format:
	isort -rc -y .
	black -l 79 .


.PHONY: lint
## lint: check everything's okay
lint:
	isort --recursive --check-only .
	black -l 79 --check .