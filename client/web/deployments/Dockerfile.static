FROM node:12.2.0-alpine AS build-stage

LABEL maintainer="Mario Raspiantoro<raspiantoro@gmail.com>"
LABEL REPO="https://github.com/raspiantoro/realtime-web-apps"

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY ./client/web .

RUN npm install

RUN npm run build


FROM nginx:1.16.0-alpine

LABEL maintainer="Mario Raspiantoro<raspiantoro@gmail.com>"
LABEL REPO="https://github.com/raspiantoro/realtime-web-apps"

COPY --from=build-stage /app/build /usr/share/nginx/html
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]