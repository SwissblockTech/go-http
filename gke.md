
# GKE

1. Spin up a GKE cluster configured as follows
	- Cluster basics
		- Regional (please use one of Europe regions)
		- Release channel: regular (please select latest K8s version)
	- Networking
		- Private cluster: true
		- Access control plane using its external IP address: true
		- Enable control plane authorized networks: false
		- `OPTIONAL` Enable Gateway API: true
	- Security
		- Enable Workload Identity
	- Features
		- Enable Cost Allocation
		- `OPTIONAL` Enable image streaming
		- Enable GKE usage metering
	- Default node pool
		- Total number of nodes: 3
		- Specify node locations: (tick all available zones)
		- Blue-green upgrade
			- Batch node count: 1
			- Batch soak duration: 300
			- Nodepool soak duration: 600
		- Nodes
			- Machine type: e2-small or e2-medium
		- Security
			- `OPTIONAL` Enable secure boot: true

2. Deploy applications
	- Deploy the HTTP server [v1.0.0](https://hub.docker.com/layers/bygui86/http-server/v1.0.0/images/sha256-e27db42723bacf94127d1d4128d6cada43f84a2e91e83dd5bf28340f591a7426?context=explore)
	- Deploy the HTTP client [v1.0.0](https://hub.docker.com/layers/bygui86/http-client/v1.0.0/images/sha256-0a69cddd13651af05af660d14f1df772f72edb6c310492dc647fbe2ad60c7e4a?context=explore)

3. Expose HTTP client over internet

4. Perform following HTTP requests
	- Create 3 new products
	- Get all products
	- Get product by ID
	- Delete product by ID
