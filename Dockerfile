# From https://docs.docker.com/reference/dockerfile/#parser-directives
# syntax=docker/dockerfile:1
# check=error=true

# Pulls base image from
# https://github.com/docker-library/postgres/blob/e00e1bd34ec5c8a8e7ad89b273b3d42efaf6d5bc/18/trixie/Dockerfile
FROM postgres:18.6

# Anything in /docker-entrypoint-initdb.d/ is executed on startup
COPY ./scripts/seed.sql /docker-entrypoint-initdb.d/

# Map local database/ to /var/lib/postgresql/ inside the container
VOLUME database:/var/lib/postgresql/

ENV POSTGRES_USER="postgres"
ENV POSTGRES_PASSWORD="secret"
