#Using container build at remote & publish
#On heroku.com (free hosting services) build docker container at remote(in heroku platform) and publish app
#create heroku app using browser login
#using heroku CLI in the root directory of the project execute following commands
#add heroku.yml file in the root with required commands
#git add heroku.yml
#git commit -m 'added heroku file'
#git push origin master
#heroku login
#heroku git:remote -a goreactcrud
#heroku buildpacks:set https://github.com/jincod/dotnetcore-buildpack (if required)
#heroku buildpacks:add --index 1 heroku/nodejs (if required)
#heroku buildpacks -- check for registered buildpacks for the repository/project
#heroku stack:set container
#git subtree push --prefix goreactcrud heroku master  OR  git push heroku master  OR  git push -f heroku master

# -- Build react client & serve static files using npm's serve package 
FROM node:11.9-alpine as clientbuilder
COPY ./client/package*.json ./client/
RUN cd ./client && npm install --silent
COPY ./client/ ./client/
RUN cd ./client && npm run build
# RUN npm config set unsafe-perm true
# RUN cd ./client && npm install -g serve
# CMD ["serve", "-s", "./client/build"]

# -- Build Go server
FROM golang:1.11.5 as serverbuilder
WORKDIR /buildapp
COPY ./server/ ./server/
COPY ./database/ ./database/

# copy react's static files from previous build i.e. clientbuilder
COPY --from=clientbuilder ./client/build ./client/build

# Required dependencies for go server project
RUN cd ./server && go get -d -v github.com/gorilla/mux github.com/jinzhu/gorm \ 
  github.com/jinzhu/gorm/dialects/sqlite \
  github.com/gorilla/handlers \
  github.com/hashicorp/go-memdb

RUN cd ./server && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main . \
  && go build ./server/*.go

# -- Build react client
# RUN curl -sL https://deb.nodesource.com/setup_10.x |  bash -
# RUN apt-get install -y nodejs
# COPY ./client/package*.json ./client/
# RUN npm install --silent --prefix ./client
# COPY ./client/ ./client/
# RUN npm run build --prefix ./client

# start the go server which would serve react's static files alongwith API endpoints
CMD ["go", "run","./server/","*.go"]
# CMD ["sh","-c","go run ./server/*.go && serve -s ./client/build"]
# CMD go run ./server/*.go ; serve -s ./client/build
# EXPOSE 8081