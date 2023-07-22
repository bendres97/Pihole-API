# Pihole-API
This repo is a work-in-progress. By no means should you deploy this to any production environment.

Contributions are always welcome.

## For Me (Because I forget things)

Generate spec from file:

```bash
docker run --rm -v $(pwd):/local openapitools/openapi-generator-cli generate -i /local/spec.yaml -g go-server -o /local/
```