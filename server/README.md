**Project Description**
This project contains a simple Go application that returns the Docker container ID and hostname on which the application is running.

**Running the Application**
To run the application as a Docker container, follow these steps:
1.Build the Docker image from the Dockerfile located in the root directory of the project:
`docker build -t name_my_project .`
2. Run the container from the built image:
`docker run -p 8080:80 name_my_project`
3. Open your browser and navigate to http://localhost:8080 to access the application.


