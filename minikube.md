
# Minikube

1. Spin up a K8s cluster using Minikube

2. Deploy applications
	- Deploy the HTTP server [v1.0.0](https://hub.docker.com/layers/bygui86/http-server/v1.0.0/images/sha256-feaab1f4839d1001635f77b78da241e005afd2b4fe519fb3c80505af9ab0a0ef?context=explore)
	- Deploy the HTTP client [v1.0.0](https://hub.docker.com/layers/bygui86/http-client/v1.0.0/images/sha256-02350f16d991394d0b453f40b09170a4be9599c2e159e7c46e925fbf035fb2bf?context=explore)

3. Expose HTTP client to localhost

4. Perform following HTTP requests
	- Create 3 new products
	- Get all products
	- Get product by ID
	- Delete product by ID
