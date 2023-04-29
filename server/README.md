**Project Description**  

This project contains a simple Go application that returns the Docker container ID and hostname on which the application is running.  

**1 - Getting started**  

1.1 Clone this repository to your local machine:  

`git clone https://github.com/AlexSHuk/dockerfiles.git`

1.2 Go to the directory of the Dockerfile you want to build:  

`cd server`  

**2 - Running the Application**  

To run the application as a Docker container, follow these steps:  

2.1 Build the Docker image from the Dockerfile located in the root directory of the project:  

`docker build -t name_my_project .`  

2.2 Run the container from the built image:  

`docker run -p 80:80 name_my_project`  

2.3 Open your browser and navigate to http://localhost:80 to access the application.  

**!!!note**: if you don't have Docker installed, use the official documentation https://docs.docker.com/engine/install/ubuntu/


