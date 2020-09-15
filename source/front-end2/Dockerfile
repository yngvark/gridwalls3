ARG NODE_TAG=12-slim

FROM node:$NODE_TAG

WORKDIR /app

COPY src ./
RUN npm install

COPY src .

RUN npm run build

CMD ["npm", "run", "start" ]
