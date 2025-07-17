# firstGoApp
# firstGoApp


## Curl Request to test the app
```bash
curl -X POST http://localhost:8080/signup -H "Content-Type: application/json" \
     -d '{"email": "test@example.com", "password": "1234"}'

curl -X POST http://localhost:8080/login -H "Content-Type: application/json" \
     -d '{"email": "test@example.com", "password": "1234"}'
```