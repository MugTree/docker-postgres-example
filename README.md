# Basic go application dockerised with a persistent postgres container

Db is created on init using the commands from the init folder

To interact with the db from the host. If you look in the docker-compose.yml you'll see that there is a port mapping that respects that postgres might be running on the host.

So to connect to to postgres container something like

            psql -h localhost -p 5433 -d mydb -U myuser -W

Should do the trick...
