# Builder image
ARG NODE_TAG=12-slim

FROM node:$NODE_TAG as builder

WORKDIR /app

COPY package* ./
RUN npm install

COPY dist/ ./dist/
COPY src/ ./src/
COPY ./tsconfig-frontend.json .
COPY rollup*.js ./

RUN npm run build

# Node, typescript stuff
COPY server/ ./server/
#COPY --chown=node ./tsconfig.json .
COPY ./tsconfig.json .
RUN npm run tsc

# Final image
FROM node:$NODE_TAG
WORKDIR /app

COPY package* ./
RUN npm install --production

#COPY node_modules ./node_modules
#COPY package* ./
COPY dist ./dist/
COPY --from=builder /app/server/*.js ./server/

CMD ["node", "server/node_server.js" ]
