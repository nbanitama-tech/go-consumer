# Go-Consumer-NSQ
Simple application on how to implement consumer with NSQ.

## Installing / Getting started
- docker-compose

## Developing

### Built With
- TBD

### Prerequisites
- docker-compose

### Setting up Dev

Here's a brief intro about what a developer must do in order to start developing
the project further:

```shell
git clone https://github.com/your/your-project.git
cd your-project/
docker-compose up
```

To test / publish message to topic, refer to part Tests

### Building
the step for building is already included on docker-compose file

### Deploying / Publishing
- TBD

## Versioning
- TBD

## Configuration
- TBD

## Tests
- To publish message to topic, use this request via postman / curl
```
curl -X POST \
  'http://localhost:4151/pub?topic=topic' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
	 "Name": "NOBA",
	 "Address": "MAIDU"
}'
```

## Style guide
- TBD

## Api Reference
- TBD

## Database
- TBD

## Licensing
- TBD