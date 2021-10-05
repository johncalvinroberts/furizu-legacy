# API Routes

### Archives

##### `GET /api/archives`
List current user's archives

##### `POST /api/archives`
Create an archive instance

##### `GET /api/archives/{id}`
Get an archive instance

##### `DELETE /api/archives/{id}`
Destroy an archive instance

---

### Whoami

##### `GET /api/whoami`
Get current user info

##### `POST /api/whoami`
Initialize whoami challenge

##### `PATCH /api/whoami/redeem`
Redeem temporary credentials for JWT

##### `PATCH /api/whoami/refresh`
Exchange JWT for fresh JWT

##### `DELETE /api/whoami/revoke`
Revoke current user's JWT
