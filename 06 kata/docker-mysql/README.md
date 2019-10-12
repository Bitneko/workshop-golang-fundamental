# Database Setup

This setup the necessary database to use with kata exercise using Docker.

Connect to the database on 127.0.0.1 on port 3306 to `todo` database using username `root` and password `root`

If you have anything listening to port 3306 in your local environment or another docker container, you need to turn those off or edit Makefile to bind 3306 to another port.

### Create the Docker image

```
make build
```

### Create the Docker container with the image and start it

```
make start
```

### Stop the container from running (Run it in a separate terminal instance)

```
make stop
```

### Stop the container and remove it

```
make clean
```