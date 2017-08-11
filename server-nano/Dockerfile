FROM node:8-alpine
ENV NODE_ENV production
WORKDIR /usr/src/app
COPY "package.json" "package-lock.json" ./
RUN npm install --silent
COPY . .
EXPOSE 3000
CMD npm start