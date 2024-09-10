# Phonebook API


### Instructions

1. **Create the `.env` file:**
   - In your project's root directory, create a new file named `.env`. 

   use this file structure:
    ```
    # .env
    DB_USER=<db_user>
    DB_PASSWORD=<db_pass>
    DB_NAME=<db_name>
    DB_HOST=<db_host>
    DB_PORT=<db_port>
    ```
    for example:
   ```
   # .env
   DB_USER=postgres
   DB_PASSWORD=pass1234
   DB_NAME=contactdb
   DB_HOST=localhost
   DB_PORT=5432
   ```
   
2. **Add your database credentials:**
   - Open the `.env` file in a text editor.
   - Replace `<db_user>`, `<db_pass>`, `<db_name>`, `<db_host>` and `<db_port>`
     with your actual database: username, password, database name, host and port.

3. **Save the file:**
   - Save and close the `.env` file.


## Run using Docker-Compose

```bash
docker-compose build --no-cache
docker-compose up
```




## API Endpoints

### 1. Get All Contacts
**GET /contacts**  
Returns a list of all contacts.

### 2. Add New Contact
**POST /contacts**  
Create a new contact.

**Payload:**
```json
{
  "first_name": "string",
  "last_name": "string",
  "phone": "string",
  "address": "string"
}
```

### 3. Update Contact
**PUT /contacts**  
Update an existing contact. Only the `id` is required, other fields can be `null` if not updating them.

**Payload:**
```json
{
  "id": "string",
  "first_name": "string or null",
  "last_name": "string or null",
  "phone": "string or null",
  "address": "string or null"
}
```

### 5. Search Contacts
**GET /contacts/search**  
Search for contacts by a search term.

**Query Parameter:**
- `term` (required): The search term to find contacts by any matching fields (first name, last name, phone, or address).

**Example:**
```
GET /contacts/search?term=john
```

**Response:**
- `200 OK`  
  Returns a list of contacts matching the search term.
  
- `400 Bad Request`  
  If the `term` parameter is missing.

**Sample Response:**
```json
[
  {
    "id": "1",
    "first_name": "John",
    "last_name": "Doe",
    "phone": "+1234567890",
    "address": "123 Main St"
  }
]
```


## Tests
- To run the tests, navigate to the `/tests` directory and execute the following command: `go test`.
