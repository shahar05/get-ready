# Phonebook API

## Run using Docker

```bash
docker build -t phonebook-api .
docker run -p 8080:8080 phonebook-api
```


### Instructions

1. **Create the `.env` file:**
   - In your project's root directory, create a new file named `.env`. 

   use this file structure:
    ```
    # .env
    DB_USER=<db_user>
    DB_PASSWORD=<db_pass>
    DB_NAME=<db_name>
    ```
2. **Add your database credentials:**
   - Open the `.env` file in a text editor.
   - Replace `<db_user>`, `<db_pass>`, and `<db_name>` with your actual database username, password, and database name.

3. **Save the file:**
   - Save and close the `.env` file.

The application will automatically read these values when it starts up. Ensure that the `.env` file is not included in version control by adding it to your `.gitignore` file.

