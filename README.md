# Basic go application dockerised with a persistent postgres container

Db is created on init using the commands from the seed-data folder and will persist over stop and starts.

To interact with the db from the host machine - If you look in the docker-compose.yml you'll see that there is a port mapping that respects that postgres might be running on the host.

So to connect to to postgres container something like

             psql -h localhost -p 5433 -d postgres -U postgres -W

Fiddled around with creating a new database and a new user in the initalisation stage and couldn't seem to get it working.
The db command kept getting run against the default database.

### Useful commands

            docker compose up -d
            docker compose down
            docker network prune

### Rebuild

            docker compose up -d --build

#### Follow logs real-time:

            docker-compose logs -f

This seems to just work without any additional commands adding to the docker compose file.
