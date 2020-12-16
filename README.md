[![Drone (cloud)](https://img.shields.io/drone/build/jones2026/go-hello-world)](https://cloud.drone.io/jones2026/go-hello-world)
[![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/jones2026/go-hello-world)](https://hub.docker.com/r/jones2026/go-hello-world/tags?page=1&ordering=last_updated)
[![Docker Pulls](https://img.shields.io/docker/pulls/jones2026/go-hello-world)](https://hub.docker.com/r/jones2026/go-hello-world)
# go-hello-world
Sample go web app to test docker deployments

## Running your Hello World docker image
Start your image:

	$ docker run -d -p 8080:8080 jones2026/go-hello-world
    3e1a2b32e73114333d362ffbb71d4bf13375f8c30e80c622540f58ebbb75ce08
	$ curl http://localhost:8080/
    Hello World!
