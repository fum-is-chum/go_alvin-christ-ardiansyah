## (21) Docker
1. Container → Abstraction at the app layer, enables us to run multiple app on the same machine, takes less space than VMs, generally faster than VMs
2. Docker → open source platform for developing, shipping and running containerized applications
3. Steps in using Docker:
   - Install docker and docker-compose
   - Create Dockerfile
     - Use small base image (alpine version) and utilize multi layer to achieve smaller final image size
   - Push your image to docker registry (optional)
   - Create a docker-compose.yml to help define and run multi-container 