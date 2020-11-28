# Builder image
ARG NODE_TAG=12-slim

FROM node:$NODE_TAG as builder

WORKDIR /app

COPY package* ./
RUN npm install

COPY dist/ ./dist/
COPY src/ ./src/
RUN ls -lrt
COPY rollup*.js ./
COPY app.js .

RUN npm run build

# Final image
FROM node:$NODE_TAG
WORKDIR /app

COPY package* ./
RUN npm install --production

#COPY node_modules ./node_modules
#COPY package* ./
COPY dist ./dist

COPY app.js .

CMD ["node", "app.js" ]
