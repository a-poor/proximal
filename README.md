# proximate

_created by Austin Poor_

Proximate is a reverse proxy written in Go.


## Sample Config File:

```yaml
---
proximate:
  upstreams:
  - name: backend
    servers:
    - localhost:8080
    - localhost:8081
    - localhost:8082
    healthcheck:
    - /health
  server:
    port: 80
    routes:
    - path: /
      static: ./static       
    - path: /api/
      forward_to: backend 
```
