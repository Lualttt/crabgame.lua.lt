# crabgame.lua.lt
---

easy way to join a crab game lobby

this is meant to be ran as a docker container behind a proxy

## docker compose.yml example
```yaml
services:
  lualt:
    image: crabgame-lualt # make sure to "docker build -t crabgame-lualt ."
    container_name: crabgame-lualt
    restart: unless-stopped
    ports:
      - 8080:8080
```
