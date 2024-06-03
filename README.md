# goexpert-rate-limiter

This project is a simple implementation of a rate limiter middleware using Go.

The middleware implemented uses a [Redis](https://redis.io/) cache in order to store the requisitions' count and blocked key status.

The configurations of the limiter can be altered by changing the variables found in the `.env` file inside the `cmd` folder.

## Commands to run the application locally

1. Just run the command below to start the redis container:

```sh
docker compose up -d
```
2. In the `http` folder there are two files with examples of requests that can be made to the server in order to test it.
If you are using VSCode as your IDE downloading the [Rest Client Plugin](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) is an excellent way execute these requests.

3. Inside the `cmd` folder you will also find end to end tests based on the configs found in the main branch of this repository.
   The tests can be ran from inside the `cmd` folder by using the following command:
```sh
go test -v
```
