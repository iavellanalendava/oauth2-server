# oauth2 server

**oAuth2 server** provides a secure and standardized approach for granting access to protected 
resources on behalf of users or applications complying with modern authentication and 
authorization systems.

---

### API
1. `POST` to `/token`: Responsible for **issuing tokens**.  
When a client presents valid credentials, such as client ID and secret under Basic Authentication 
authorization method, the server verifies them and generates an access token. This token is then 
returned to the client, enabling subsequent authenticated requests to the protected resources.


2. `POST` to `/verify`: Designed for **token introspection**.  
Clients can send an access token to this endpoint to determine its validity and 
obtain metadata about it. The server performs an integrity check on the token and provides 
information such as the associated user or client, expiration time, and the scope of access granted.

3. `POST` to `/keys`: Designed for **listing all signing keys** for a specific user.
In OAuth2, tokens are often digitally signed to ensure their authenticity. 
This endpoint allows clients to retrieve a list of public signing keys that can be used to 
verify the integrity and authenticity of issued tokens.  

---

### Pending
- [ ] Deployment manifests to deploy the server in Kubernetes cluster
- [ ] Storage of credentials and keys in an appropriate database
- [ ] Addition of logs
