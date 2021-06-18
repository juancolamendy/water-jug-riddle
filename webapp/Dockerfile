FROM mhart/alpine-node:12
MAINTAINER JC

RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN rm -fr node_modules && npm install && npm run build

EXPOSE 3000

CMD ["npm", "run", "start"]