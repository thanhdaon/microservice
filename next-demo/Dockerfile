FROM mhart/alpine-node:12.19.0
WORKDIR /app
COPY src/package.json /app/package.json
COPY src/yarn.lock /app/yarn.lock
RUN yarn install
COPY src .
RUN yarn build

FROM mhart/alpine-node:12.19.0
WORKDIR /app
COPY src/package.json /app/package.json
COPY src/yarn.lock /app/yarn.lock
ENV NODE_ENV=production
RUN yarn install --prod

FROM mhart/alpine-node:slim-12.19.0
WORKDIR /app
COPY --from=0 /app/.next /app/.next
COPY --from=1 /app/node_modules /app/node_modules
COPY src/public /app/public
CMD ["/app/node_modules/.bin/next", "start"]
