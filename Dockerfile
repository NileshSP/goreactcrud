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
FROM golang:1.11.5 as builder
WORKDIR /buildapp
COPY ./server/ ./server/
COPY ./database/ ./database/
RUN cd ./server && go get -d -v github.com/gorilla/mux github.com/jinzhu/gorm github.com/jinzhu/gorm/dialects/sqlite github.com/gorilla/handlers
RUN cd ./server && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN go build ./server/*.go
CMD ["go", "run","./server/","*.go"]
#EXPOSE 8081

# -- 2nd Step :- Build react client
# FROM node:11.9-alpine as publishbuilder
# COPY ./client/package*.json ./client/
# RUN cd ./client && npm install --silent
# COPY ./client/ ./client/
# RUN cd ./client && npm run build
# RUN npm config set unsafe-perm true
# RUN cd ./client && npm install -g serve
# CMD ["serve", "-s", "./client/build"]
# EXPOSE 8081

# FROM nginx
# COPY --from=publishbuilder ./client/build /usr/share/nginx/
# COPY ./client/nginx.conf /etc/nginx/nginx.conf
# EXPOSE 8081
# ENTRYPOINT ["nginx","-g","daemon off;"]