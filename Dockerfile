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


# -- 1st Step :- Build Go server
# Get the Go with version as specified
FROM golang:1.11.5 as builder

# Set the working directory to the buildapp directory
WORKDIR /buildapp

# Copy all contents from the root
COPY ./server/ ./server/

COPY ./database/ ./database/

RUN cd ./server && go get -d -v github.com/gorilla/mux github.com/jinzhu/gorm github.com/jinzhu/gorm/dialects/sqlite github.com/gorilla/handlers

RUN cd ./server && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Build the go project in linux server environment
RUN go build ./server/*.go

# Start the server to listen for requests
CMD ["go", "run", "./server/server.go"]
EXPOSE 8081

# -- 2nd Step :- Build react client
# You should always specify a full version here to ensure all of your developers
# are running the same version of Node.
FROM node:11.8.0

# Override the base log level (info).
ENV NPM_CONFIG_LOGLEVEL warn

# Install and configure `serve`.
RUN npm install -g serve
CMD serve -s build
EXPOSE 3000

# Install all dependencies of the current project.
COPY ./client/package*.json ./client/
#COPY npm-shrinkwrap.json npm-shrinkwrap.json
RUN cd ./client && npm install 

# Build for production.
RUN cd ./client && npm run build --production