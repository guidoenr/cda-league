# Cda-League
> :bulb: this was never deployed to any environment, development stopped when client was unable to pay.


# Local 
### connect to db in linx
`sudo -u postgres psql postgres`

### creating the database
`CREATE DATABASE fulbo`

### creating an user
`CREATE user root WITH encrypted password 'root';`

### granting permissions
`GRANT ALL privileges ON database fulbo to root;`

### USERS
```bash
user=root
password=root 
db=fulbo
```
# Deploying
## use buildpacks:
1. `go` buildpack
2. `react` buildpack

## see current buildpacks
- `heroku buildpacks`

## add buildpacks [CLI]
- `heroku buildpacks:add --index 2 https://github.com/mars/create-react-app-buildpack.git`
- the go buildpack is detected automatically
