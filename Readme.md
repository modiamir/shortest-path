Run below command to start the app:
```bash
docker-compose up -d
```

And send below request to get the shortest path between two point with maximum of 4 edges:
```http request
GET http://127.0.0.1:8080/v1/shortest-path?from=LYR&to=MEL
```