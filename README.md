# book-svc

Service for storing books information with basic CRUD operations on the API.

To create a book, the frontend side gets a signed EIP712 create message from the backend when sending a POST method. Simultaneously, a “raw” book (with basic information) is being added to the database. Then `update_listener` on the tracker service will update mocked fields when the contract is deployed. To link the contract with our raw book, we use `token_id` field.

