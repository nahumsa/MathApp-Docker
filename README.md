# Web app using Golang and Docker

[![MathApp](https://circleci.com/gh/nahumsa/MathApp-Docker.svg?style=svg)](https://circleci.com/gh/nahumsa/MathApp-Docker)

This project is based on the post by [SemaphoreCI](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker) on using Golang and Docker for building a web app. 

I chose to use [Gin Gonic](https://github.com/gin-gonic/gin) as the web framework for this simple application and added tests for the http handlers.

I used circleCI for Continous Integration and Github Actions for Continous Deployment of the Docker image on Heroku.

In order to run the docker container on your local machine, use the following commands:

`docker build -t go-mathapp`

`docker run -d -p 8080:8080 go-mathapp`

Or simply run `rundocker.sh`.


# Deploying on heroku

In order to deploy on heroku manually run the following commands:

- `heroku create go-mathapp`
- `heroku login`
- `heroku container:login`
- `heroku container:push web -a go-mathapp`
- `heroku container:release web -a go-mathapp`

You can see the deployed website on [heroku](https://go-mathapp.herokuapp.com/)