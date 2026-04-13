# Testing Endpoint


### Health Check 
- Endpoint
```bash
GET

{{BACKEND_BASE_URL}}/v1/health

```

### Example

<img width="1863" height="787" alt="image" src="https://github.com/user-attachments/assets/4b3facd6-b2e8-4a75-a977-82d7edfc0361" />

---

### Login 


- Endpoint
```bash
POST

https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key={{FIREBASE_API_KEY}}
```

### Example

<img width="1906" height="926" alt="image" src="https://github.com/user-attachments/assets/cb662d22-9918-48c1-af16-0314160e598a" />

### Verifiy Token 

- Endpoint
```bash
POST

{{BACKEND_BASE_URL}}/v1/auth/verify-token
```
### Example
<img width="1862" height="980" alt="image" src="https://github.com/user-attachments/assets/2379f67d-e367-45d0-ba13-103943eede01" />
<img width="1899" height="924" alt="image" src="https://github.com/user-attachments/assets/964a93e1-3c63-4b79-a7bc-10be0199407d" />



---

### Get All Product 
- Endpoint
```bash
GET

{{BACKEND_BASE_URL}}/v1/products
```

### Example
<img width="1911" height="884" alt="image" src="https://github.com/user-attachments/assets/3f3a0148-5989-458d-8a22-a2c262011475" />


---

### Get Product By id 
- Endpoint
```bash
GET

{{BACKEND_BASE_URL}}/v1/products/2
```

### Example
<img width="1914" height="921" alt="image" src="https://github.com/user-attachments/assets/a52e91a4-db05-4a8d-b7fe-b1a3036d92d2" />

---

# Seeder

- Command Running Seed
```bash
go run seeds/seed.go
```

### Setelah di run, akan tampil seperti di gambar berikut
<img width="1889" height="869" alt="image" src="https://github.com/user-attachments/assets/499e529c-e106-4b31-a407-79de149fd190" />


---

# Flow (Sequence Diagram)

### Create User
<img width="6250" height="3165" alt="User Creation Request Flow-2026-04-11-085325" src="https://github.com/user-attachments/assets/755e8604-7da3-4ade-b323-fade9615863e" />

### Code
```bash
sequenceDiagram
participant Client
participant Middleware
participant Handler
participant Service
participant Repository
participant DB
Client->>Middleware: HTTP Request
Middleware->>Handler: forward request
Handler->>Service: CreateUser()
Service->>Repository: SaveUser()
Repository->>DB: INSERT USER
DB-->>Repository: OK
Repository-->>Service: user saved
Service-->>Handler: user response
Handler-->>Client: JSON Response
```

# Get All Product
<img width="2908" height="1206" alt="mermaid-diagram-2026-04-13-232415" src="https://github.com/user-attachments/assets/5d0e3889-cafa-403a-9d34-de2444ef1f34" />

### Code

```bash
sequenceDiagram
participant Client
participant Middleware
participant Handler
participant Service
participant Repository
participant DB
Client->>Middleware: HTTP Request (Get v1/products)
Middleware->>Handler: forward request
Handler->>Service: GetAll
Service->>Repository: FindAll
Repository->>DB: SELECT * FROM products
DB-->>Repository: product rows
Repository-->>Service: products list
Service-->>Handler: products response
Handler-->>Client: JSON Response
```

# Get Product by Id
<img width="7845" height="3165" alt="Product Retrieval Flow-2026-04-13-163347" src="https://github.com/user-attachments/assets/335ca494-570b-43fa-854a-a26caa263de7" />

### Code

```bash
sequenceDiagram
participant Client
participant Middleware
participant Handler
participant Service
participant Repository
participant DB

Client->>Middleware: HTTP Request (Get v1/products/id)
Middleware->>Handler: forward request
Handler->>Service: GetById
Service->>Repository: FindById
Repository->>DB: SELECT * FROM products where id=?
DB-->>Repository: product rows
Repository-->>Service: products list
Service-->>Handler: products response
Handler-->>Client: JSON Response
```



