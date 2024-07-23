
# HTTP server

## Prerequisites

`/!\ WARN` Requires PostgreSQL database

## Ports

- API: `8080`
- metrics: `9090`

## API endpoints

| Method | URL            | Description                                |
|--------|----------------|--------------------------------------------|
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
| DB_HOST           | PostgreSQL host          |          | localhost |                                         |
| DB_PORT           | PostgreSQL port          |          | 8080      |                                         |
| DB_USERNAME       | PostgreSQL username      | *        | 8080      |                                         |
| DB_PASSWORD       | PostgreSQL password      | *        | 8080      |                                         |
| DB_NAME           | PostgreSQL DB name       | *        | 8080      |                                         |
| DB_SSL_MODE       | PostgreSQL SSL mode      |          | 8080      |                                         |
| REST_HOST         | API host                 |          | localhost |                                         |
| REST_PORT         | API port                 |          | 8080      |                                         |
