# CircleCI Auditr

Audit circleci projects for some things we care about.

## Usage

```bash
go build .
./circleci-audit -help
```

## Configuration

The application expects to find secrets in environment variables:
- CIRCLE_TOKEN

It is recommended to use [direnv](https://direnv.net/) to manage these.

### Create a CircleCI token

1. Go to circleci.com/account/api. Sign in with Github if necessary.
2. Create an API token.
3. Save the token as CIRCLE_TOKEN in `.envrc`.

