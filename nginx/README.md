**Project Description**  

This directory contains a Dockerfile for building a container image that runs Nginx web server with Alpine OS;  

**1 - Getting started**  

1.1 Clone this repository to your local machine:

`git clone https://github.com/AlexSHuk/dockerfiles.git`  

1.2 Go to the directory of the Dockerfile you want to build:

`cd nginx`  

**2 - Running the Application**  

To run the application as a Docker container, follow these steps:

2.1 Build the Docker image from the Dockerfile located in the root directory of the project:

`docker build -t name_my_project .`

2.2 Run the container from the built image:

`docker run -p 80:80 name_my_project`

2.3 Open your browser and navigate to http://localhost:80 to access the application.  

**NOTE** 
the commented lines contain the user's nginx configuration file and the resource directory.  
Create your `nginx.conf` configuration file and uncomment the lines.  
If you are using docker-compose you can specify the path to the configuration in docker-compose.yml, for example:  
`volumes:`
`- ./path-to/nginx.conf:/etc/nginx/nginx.conf`
`- ./path-to/source/:/usr/share/nginx/html/`
 



