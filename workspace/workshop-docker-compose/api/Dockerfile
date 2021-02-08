FROM node:10.16.0-alpine AS dependencies

ENV TERM=xterm
ENV TZ=Asia/Bangkok

RUN mkdir -p /usr/app
WORKDIR /usr/app

COPY ./package*.json ./

RUN npm install --quiet

COPY . .
