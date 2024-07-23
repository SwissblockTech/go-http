
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
	- Deploy the HTTP server [v1.0.0](https://hub.docker.com/layers/bygui86/http-server/v1.0.0/images/sha256-feaab1f4839d1001635f77b78da241e005afd2b4fe519fb3c80505af9ab0a0ef?context=explore)
	- Deploy the HTTP client [v1.0.0](https://hub.docker.com/layers/bygui86/http-client/v1.0.0/images/sha256-02350f16d991394d0b453f40b09170a4be9599c2e159e7c46e925fbf035fb2bf?context=explore)

3. Expose HTTP client over internet

4. Perform following HTTP requests
	- Create 3 new products
	- Get all products
	- Get product by ID
	- Delete product by ID
