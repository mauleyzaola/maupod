build: install
	npm run build

browser: install
	PORT=9990 npm run start

dev: install
	# this value should be in the environment instead
	# REACT_APP_API_URL="http://localhost:9999" npm run start
	BROWSER=none PORT=9990 npm run start

install:
	yarn install
