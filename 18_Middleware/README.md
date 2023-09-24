## (18) Middleware
1. Middlewarer is an entity that hooks into server's request/response processing, which allow us to implement various functionalities around incoming http requests to your server and outgoing responses
2. Some example of middlewares:
   - Basic Auth
   - JWT
   - Logger
   - Request Limiter
   - Proxy
   - CORS
   - Session
   - Trailing slash
   - Rewrite https, etc.

3. In Echo, there is two types of middleware:
   - Echo#Pre => executed before router processes the request
   - Echo#Use => executed after router processes the request and has full access to echo.Context API