# Install deps
FROM node:10 AS base
WORKDIR /base
COPY . .
RUN yarn install

# Building app
FROM base AS build
ENV NODE_ENV=production
WORKDIR /build
COPY --from=base /base ./
RUN yarn run build

# Create production container
FROM node:10 AS production
ENV NODE_ENV=production
WORKDIR /app
COPY --from=build /build/package.json ./
COPY --from=build /build/yarn.lock ./
COPY --from=build /build/.next ./.next
COPY --from=build /build/public ./public
RUN yarn add next --frozen-lockfile --production && yarn cache clean

# Running the app
EXPOSE 3000
CMD [ "yarn", "start" ]
