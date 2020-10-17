FROM node:12

WORKDIR /usr/src/app

# Install dependencies
COPY package*.json ./
RUN npm install

COPY main.js .

EXPOSE 3000
HEALTHCHECK --interval=20s --timeout=30s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:3000/_healthcheck || exit 1

CMD [ "node", "main.js" ]