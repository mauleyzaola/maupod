FROM node:10.20.1-alpine3.9
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install