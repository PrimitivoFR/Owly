# Stage1 BUILDER
FROM node:10-alpine as builder

# Copy the package.json to install dependencies
COPY package.json package-lock.json ./

# Install the dependencies and init the folder
RUN npm install && mkdir /owly-client && mv ./node_modules ./owly-client

WORKDIR /owly-client

COPY . .

# Build the project and copy the files
RUN npm run build:prod

# Stage 2 DEPLOY
FROM nginx:alpine

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf

# Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*

# Copy from the Stage1
COPY --from=builder /owly-client/dist /usr/share/nginx/html

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]