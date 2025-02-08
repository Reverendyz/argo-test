FROM node:23-alpine as builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm test

COPY . .

RUN npm run build

ARG HOST
ARG PORT

RUN VITE_API_HOST=$HOST VITE_API_PORT=$PORT npm run build

FROM nginx:1.9.15-alpine

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]


