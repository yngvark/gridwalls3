# Builder image
ARG NODE_TAG=12-slim

FROM node:$NODE_TAG as builder

WORKDIR /app

COPY src/package* ./
RUN npm install --production

COPY src ./

RUN npm run build

# Final image
FROM node:$NODE_TAG
WORKDIR /app

COPY src/node_modules ./node_modules
COPY src/package* ./
COPY src/dist ./dist
COPY src/public ./public

COPY src/app.js .

CMD ["node", "app.js" ]
