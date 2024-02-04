# Header-to-QueryParameter Middleware
Inverse of forked repo.

Example Traefik config:
``` yaml
[...]
experimental:
  plugins:
    queryParameterToHeader:
      moduleName: github.com/corticph/queryparameter-to-header
      version: v1.0.0

http:
[...]
  routers:
    whoami:
      rule: "Host(`whoami.localhost`)"
      service: "whoami"
      entryPoints:
        - web
      middlewares:
        - add-header-from-query
  
  middlewares:
    add-header-from-query:
      plugin:
        queryParameterToHeader:
          queryParameter: "v"
          header: "X-Version"
```
