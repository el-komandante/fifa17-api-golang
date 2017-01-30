FROM scratch
ADD main /
# ENV PG_USER postgres
# ENV PG_PASSWORD postgres
EXPOSE 3002
CMD ["/main"]
