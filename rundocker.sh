# Build docker container
sudo docker build -t go-mathapp .

# Execute container
sudo docker run -d -p 8080:8080 go-mathapp