# docker-compose-utils
some useful command line tolls for docker compose file processing


## Usage

copy all environment variable from .env file into docker-compose.yml  

```
./dc-utils freeze -i docker-compose.yml -e .env -o test.yml
```