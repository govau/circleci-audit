# CircleCI Auditr

Audit circleci projects for some things we care about.

## Usage

```bash
go build .
./circleci-audit -help
```

Example output:

```
+-------------------------------------------------------------+-------+--------------+
|                             URL                             |  OSS  | HAS ENV VARS |
+-------------------------------------------------------------+-------+--------------+
| https://circleci.com/gh/govau/cga-docs                      | false | true         |
-------------------------------------------------------------+-------+--------------+
```

## Configuration

The application expects to find secrets in environment variables:
- CIRCLE_TOKEN
- GITHUB_TOKEN 

It is recommended to use [direnv](https://direnv.net/) to manage these.

### Create a CircleCI token

1. Go to circleci.com/account/api. Sign in with Github if necessary.
2. Create an API token.
3. Save the token as CIRCLE_TOKEN in `.envrc`.

### Create a Github personal access token

1. Login to github as an appropriate machine user with admin access to the required github orgs.
2. Go to https://github.com/settings/tokens
3. Create a Personal Access Token with `repo:public_repo` scope.
4. Save the token as GITHUB_TOKEN in `.envrc`.
