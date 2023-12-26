# Pihole-API
This repo is a work-in-progress. By no means should you deploy this to any production environment.

Contributions are always welcome.

## For Me (Because I forget things)

Generate spec from file:

```bash
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli generate -i /local/spec.yaml -g go-server -o /local/
```

# Actions Statuses
[![CodeQL](https://github.com/bendres97/Pihole-API/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/bendres97/Pihole-API/actions/workflows/codeql-analysis.yml)

[![Go](https://github.com/bendres97/Pihole-API/actions/workflows/go.yml/badge.svg)](https://github.com/bendres97/Pihole-API/actions/workflows/go.yml)

[![OpenAPI Spec Generator](https://github.com/bendres97/Pihole-API/actions/workflows/spec.yml/badge.svg)](https://github.com/bendres97/Pihole-API/actions/workflows/spec.yml)

[![Release](https://github.com/bendres97/Pihole-API/actions/workflows/release.yml/badge.svg)](https://github.com/bendres97/Pihole-API/actions/workflows/release.yml)