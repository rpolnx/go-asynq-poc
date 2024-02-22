Proof of concept of asynq

## Introduction

This repository means to test the project [Go hibiken/asynq](https://github.com/hibiken/asynq)

## Diagram for common jobs

![Queue graphs](./files/queue_graphs.png)

![Go hibiken/asynq](./files/diagram_background_processor.jpg)


## Scheduled jobs

![Scheduled jobs](./files/scheduled_jobs.png)


## Tasks processing

![Tasks processing](./files/jobs_processing.png)




## Running

Run dependencies
```sh
docker compose up -d
```

- Needs to configure .env first similiar to .env.example
```
    go run main.go
```

### Default host & port to access the app

- http://localhost:8080 # web server
- http://localhost:5040 # work UI
