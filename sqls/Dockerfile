FROM postgres:10.10
COPY ./init.sql /docker-entrypoint-initdb.d/
#ENV POSTGRES_PASSWORD secret
#ENV POSTGRES_USER user1
#ENV POSTGRES_DB game
EXPOSE 5032