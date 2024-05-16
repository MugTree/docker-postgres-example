# Basic go application dockerised with a persistent postgres container

**! docker compose not yet tested**

Db is created on init using the commands from the init folder and will persist over stop and starts..

To interact with the db from the host. If you look in the docker-compose.yml you'll see that there is a port mapping that respects that postgres might be running on the host.

So to connect to to postgres container something like

            psql -h localhost -p 5433 -d mydb -U myuser -W

### Useful commands

docker compose up -d  
docker compose down
docker network prune
docker run -dpr 8000:8000 postgres-test-app
docker build -t postgres-test-app .
