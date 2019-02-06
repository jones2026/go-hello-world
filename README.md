# go-hello-world
Sample go web app to test docker deployments

## Usage
To create the image execute the following command on the go-hello-world folder:

	$ docker build -t jones2026/go-hello-world .

You can now push your new image to the registry:

	$ docker push jones2026/go-hello-world

## Running your Hello World docker image
Start your image:

	$ docker run -d -p 80:80 jones2026/go-hello-world
    3e1a2b32e73114333d362ffbb71d4bf13375f8c30e80c622540f58ebbb75ce08
	$ curl http://localhost/
    The current machine timestamp in UTC: 2015-10-21T00:00:00.000000+00:00



## Endpoints

|Endpoint|Parameters|Response|Description|
|:-----:|:-----:|:----|:----------|
|/||200|Replies with current timestamp in UTC format
|/hello|name=${INSERT_NAME}|200|Will prompt with an formatted HTML page greeting using a non-embedded CSS file
|/healthz|error=${INSERT_ERROR}|200/500|replies with 500 error if error is provided, otherwise request succeeded
|/metrics||200|Prometheus metrics endpoint
