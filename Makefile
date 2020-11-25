test-all: test-backend

test-backend:
	$(MAKE) -C ./src test

test-ui:
	$(MAKE) -C ./ui-player test
