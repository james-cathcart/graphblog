FROM ubuntu

COPY bin/api-server .

ENV GB_HOST="localhost"
ENV GB_PORT="8080"
ENV GB_DB_HOST="graphblog-db"
ENV GB_DB_PORT="5432"
ENV GB_DB_USER="blog_user"
ENV GB_DB_PASS="devpass"

EXPOSE 8080

ENTRYPOINT ["./api-server"]