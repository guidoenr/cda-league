# fulbo
1978 - 1986 - 2022

# connect to db in linx
`sudo -u postgres psql postgres`

# creating the database
`CREATE DATABASE fulbo`

# creating an user
`CREATE user root WITH encrypted password 'root';`

# granting permissions
`GRANT ALL privileges ON database fulbo to root;`

## USERS
```bash
user=root
password=root 
db=fulbo
```