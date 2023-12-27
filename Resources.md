
## Resources

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

make sure to export go path

- [Swagger](https://github.com/swaggo/swag#getting-started)
- [Swagger for fiber](https://github.com/arsmn/fiber-swagger)
- [Go fiber routes](https://dev.to/koddr/go-fiber-by-examples-delving-into-built-in-functions-1p3k#simple-routes)

Generate swagger docs
```bash
swag init --parseDependency --parseInternal
```


Installing postgres
https://www.digitalocean.com/community/tutorials/how-to-install-postgresql-on-ubuntu-20-04-quickstart


CREATE USER "500BookClubU" WITH ENCRYPTED PASSWORD '500BookClubU';
CREATE DATABASE "500BookClubD" OWNER "500BookClubU";


 GRANT ALL ON DATABASE  "500BookClubD" to "postgres";

 GRANT CONNECT ON DATABASE "500BookClubD" TO "postgres";



sudo -i -u postgres

select * from "activation_codes";

3063ace1-94e4-4002-a603-6600a86ebfef