FROM node:16-buster-slim as development

WORKDIR /usr/src/app

ENV NPM_CONFIG_PREFIX=/home/node/.npm-global
ENV PATH=$PATH:/home/node/.npm-global/bin

COPY ./package.json .
COPY ./package-lock.json .

RUN npm install

COPY . .

#CMD [ "npm", "run", "dev" ]