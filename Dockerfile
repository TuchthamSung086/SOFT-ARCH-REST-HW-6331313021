# Use the official PostgreSQL image from Docker Hub
FROM postgres:latest

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=mydatabase

# Expose the PostgreSQL default port
EXPOSE 5432

# Optionally, you can run custom SQL scripts during container initialization
# Place your SQL scripts in the /docker-entrypoint-initdb.d/ directory
# COPY ./init.sql /docker-entrypoint-initdb.d/
