
# HTTP client

## Ports

- API: `8080`
- metrics: `9090`

## API endpoints

| Method | URL            | Description                                |
|:-------|:---------------|:-------------------------------------------|
| GET    | /products      | Fetch list of products                     |
| GET    | /products/{id} | Fetch a product by ID                      |
| POST   | /products      | Create a new product                       |
| PUT    | /products/{id} | Update an existing product retrieved by ID |
| DELETE | /products/{id} | Delete a product by ID                     |

## Payloads

### Product

```json
{
  "id": "string",
  "name": "string",
  "price": "float64"
}
```

## Environment variables

| Name              | Description              | Required | Default   | Available values                        |
|:------------------|:-------------------------|:---------|:----------|:----------------------------------------|
| LOG_ENCODING      | Logging encoding         |          | console   | json, console                           |
| LOG_LEVEL         | Logging level            |          | debug     | trace, debug, info, warn, error , fatal |
| ENABLE_MONITORING | Enable monitoring toggle |          | false     | true, false                             |
| MONITOR_HOST      | Monitoring API host      |          | localhost |                                         |
| MONITOR_PORT      | Monitoring API port      |          | 9090      |                                         |
| REST_SERVER_HOST  | HTTP server API host     |          | localhost |                                         |
| REST_SERVER_PORT  | HTTP server API port     |          | 8080      |                                         |
| REST_HOST         | API host                 |          | localhost |                                         |
| REST_PORT         | API port                 |          | 8080      |                                         |
