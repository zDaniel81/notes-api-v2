# API Documentation

This document provides an overview of the routes and endpoints available in the API.

### Base URL

The base URL for all endpoints is: http://localhost:8080

## Note Endpoints

### Get a Note
Retrieves a specific note by its ID.

Route: `/note/{note_id}`

Method: `GET`

#### Parameters:

note_id (path parameter) - The ID of the note to retrieve.

Response:

**200 OK** - Returns the note with the specified ID.
**404 Not Found** - If the note with the specified ID does not exist.

---

### Get All Notes
Retrieves all notes.

Route: `/notes`

Method: `GET`

Response:

**200 OK** - Returns a list of all notes.

---

### Delete a Note
Deletes a specific note by its ID.

Route: `/note/{note_id}`

Method: `DELETE`

Parameters:

note_id (path parameter) - The ID of the note to delete.
Response:

**200 OK** - Returns the deleted note with the specified ID.
**404 Not Found** - If the note with the specified ID does not exist.

---

### Create a Note
Creates a new note with the specified title and content.

Route: `/note/{title}/{content}`

Method: `POST`

Parameters:

**title (path parameter)** - The title of the note.
**content (path parameter)** - The content of the note.
Response:

**200 OK** - If the note is successfully created.

---

### Update Note Content
Updates the content of a specific note by its ID.

Route: `/note/{note_id}/content/{content}`

Method: `PUT`

Parameters:

**note_id (path parameter)** - The ID of the note to update.
**content (path parameter)** - The new content for the note.
Response:

**200 OK** - If the note is successfully updated.
**404 Not Found** - If the note with the specified ID does not exist.

---

### Update Note Title
Updates the title of a specific note by its ID.

Route: `/note/{note_id}/title/{title}`

Method: `PUT`

Parameters:

**note_id (path parameter)** - The ID of the note to update.
**title (path parameter)** - The new title for the note.
Response:

**200 OK** - If the note is successfully updated.
**404 Not Found** - If the note with the specified ID does not exist.
