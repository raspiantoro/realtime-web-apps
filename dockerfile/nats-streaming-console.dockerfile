FROM node:8-alpine

ENV CODE /usr/src/app
WORKDIR $CODE

RUN apk add git && apk add --update make

RUN git clone https://github.com/KualiCo/nats-streaming-console.git

WORKDIR $CODE/nats-streaming-console

RUN yarn

RUN yarn build-css
RUN yarn build

EXPOSE 8282
CMD ["node", "server"]