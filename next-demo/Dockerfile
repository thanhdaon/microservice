FROM mhart/alpine-node:12.19.0
WORKDIR /app
COPY package.json /app/package.json
COPY yarn.lock /app/yarn.lock
RUN yarn install
COPY . .
RUN yarn build

FROM mhart/alpine-node:12.19.0
WORKDIR /app
COPY package.json /app/package.json
RUN yarn install --production

FROM mhart/alpine-node:slim-12.19.0
WORKDIR /app
COPY --from=0 /app/.next /app/.next
COPY --from=1 /app/node_modules /app/node_modules
COPY public /app/public
CMD ["/app/node_modules/.bin/next", "start"]
