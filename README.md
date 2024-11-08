# URL Shortener

## What is this?

A URL shortener made with Golang, PostgreSQL, and Next.js

## How to run

To run the project, you need to have Docker installed.

```bash
make run
```

You can stop all the containers/services with:

```bash
make stop
```

If you need to re-build the containers/services, you can use:

```bash
make build
```

## Tests

To run the tests, you need Go installed:

```bash
make test
```

## TODO

- [ ] Add login/register via Clerk
- [ ] Only logged in users can shorten URLS
- [ ] Add counter for clicks
- [ ] Add a way to delete URLs
- [ ] Set TTL for URLs