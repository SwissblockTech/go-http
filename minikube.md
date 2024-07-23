
# Minikube

1. Spin up a K8s cluster using Minikube

2. Deploy applications
	- Deploy the HTTP server [v1.0.0](https://hub.docker.com/layers/bygui86/http-server/v1.0.0/images/sha256-e27db42723bacf94127d1d4128d6cada43f84a2e91e83dd5bf28340f591a7426?context=explore)
	- Deploy the HTTP client [v1.0.0](https://hub.docker.com/layers/bygui86/http-client/v1.0.0/images/sha256-0a69cddd13651af05af660d14f1df772f72edb6c310492dc647fbe2ad60c7e4a?context=explore)

3. Expose HTTP client to localhost

4. Perform following HTTP requests
	- Create 3 new products
	- Get all products
	- Get product by ID
	- Delete product by ID
