FROM node:10
WORKDIR /app
COPY ./cmd/socket .
RUN yarn install
CMD ["node", "server.js"]
EXPOSE 8080