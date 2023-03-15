FROM golang:alpine AS buildmaster
WORKDIR /go/src/bitbucket.org/isbtotogroup/isbpanel_backend
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# ---- Svelte Base ----
FROM node:lts-alpine AS sveltebasemaster
WORKDIR /svelteapp
COPY [ "sveltemdb/package.json" , "sveltemdb/yarn.lock" , "sveltemdb/rollup.config.js" , "./"]

# ---- Svelte Dependencies ----
FROM sveltebasemaster AS isbpsveltedepmaster
RUN yarn
RUN cp -R node_modules prod_node_modules

#
# ---- Svelte Builder ----
FROM sveltebasemaster AS sveltebuildermaster
COPY --from=isbpsveltedepmaster /svelteapp/prod_node_modules ./node_modules
COPY ./sveltemdb .
RUN yarn build

# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest as totosvelterelease
WORKDIR /app
RUN apk add tzdata
RUN mkdir -p ./sveltemdb/public
COPY --from=sveltebuildermaster /svelteapp/public ./sveltemdb/public
COPY --from=buildmaster /go/src/bitbucket.org/isbtotogroup/isbpanel_backend/app .
COPY --from=buildmaster /go/src/bitbucket.org/isbtotogroup/isbpanel_backend/env-sample /app/.env

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 9091
CMD ["./app"]