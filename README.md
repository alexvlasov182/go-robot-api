# go-robot-api

## Overview

go-robot-api is a RESTful API built using Golang, leveraging the Gin web framework, PostgreSQL for data storage, and Docker for containerization. This project showcases my experience with building scalable and efficient backend services, handling CRUD operations, and integrating a relational database.

## Key Features

- RESTful API: Implements full CRUD operations for managing robot data.
- Gin Web Framework: Provides a fast and lightweight framework for building web applications.
- PostgreSQL: Utilized as the database for robust data management.
- Docker Integration: Ensures a consistent and reproducible environment for development and deployment.
- Modular Code Structure: Maintains clean and maintainable code for scalability.

## Technologies Used

- Golang: Core language for building the API.
- Gin: A fast and simple web framework for HTTP routing and middleware.
- PostgreSQL: Relational database for persistent data storage.
- Docker: Used for containerizing the application and database.
- SQL Drivers: For connecting and interacting with PostgreSQL.

## Setup and Installation

1. Clone this repository

```
git clone https://github.com/alexvlasov182/go-robot-api.git
cd go-robot-api
```

2. Run with Docker:

- Ensure Docker is installed and running.
- Build and start the containers:

```
docker-compose up --build
```

3. Access the API:

- The API will be available at http://localhost:3000 (default port)

## How It Works

- Gin Framework: Manages routing, middleware, and response handling.
- PostgreSQL Integration: Stores and retrieves robot data using SQL queries.
- Docker: Provides an isolated environment for the app and database to run seamlessly.
- Environment Variables: Configured via a .env file for secure and flexible setup.

## Future Enhancements

- Add JWT authentication for secure endpoints.
- Implement additional filters and query parameters for robot data retrieval.
- Integrate Swagger for API documentation.

## License

This project is open-source and available under the MIT License.

## Contact

For questions or collaboration:

- Email: drumlife182@gmail.com
- LinkedIn: [LinkedIn Profile](https://www.linkedin.com/in/oleksandr-vlasov-9969ab19b/)
