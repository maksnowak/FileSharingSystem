# Dockerfile example from https://v2.vuejs.org/v2/cookbook/dockerize-vuejs-app.html?redirect=true
# Will be modified to fit the project in the future

# build stage
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]