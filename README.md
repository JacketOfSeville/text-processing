## Text Processing

Is a project with the goal of studying systems arch, so it will be overly complicated.

The idea of it is to be a microkernel at the center, that will recieve text data and pass it down to its plugins do whatever they want, such as stats collection

### How to run

This project uses docker & docker-compose as environments, so you will need them in order to run the project as intended.

[Install Docker](https://docs.docker.com/get-docker/)
[Install Docker compose](https://docs.docker.com/compose/install/)

To run the project you can just use the following command

```bash
$ docker-compose up
```

And to stop all containers you can just

```bash
$ docker-compose down
```

To force a rebuild you can:

```bash
$ docker-compose build --no-cache && docker-compose up
```

### Architecture

![alt text](http://github.com/Gustrb/text-processing/blob/main/architecture.png?raw=true)
