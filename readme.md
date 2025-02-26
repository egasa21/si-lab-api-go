# Practicum Platform API

This is a REST API for the Practicum platform, built using Golang and PostgreSQL.

## Tech Stack
- **Golang**: Backend development
- **PostgreSQL**: Database
- **Postman**: API testing and documentation

## Setup

### Prerequisites
Make sure you have the following installed:
- [Golang](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Postman](https://www.postman.com/) (optional for testing)

### Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/egasa21/si-lab-api-go.git
   cd si-lab-api-go
   ```
2. Set up environment variables:
   ```sh
    DB_HOST=your_host
    DB_PORT=your_port
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    APP_PORT=your_app_port
    LOG_LEVEL=debug
    LOG_ERROR_STACK=true
   ```
3. Start the server:
   ```sh
   go run ./cmd/api
   ```

## API Documentation
You can access the API documentation and test endpoints using Postman:

[Postman Documentation](https://documenter.getpostman.com/view/10017926/2sAYdfpWHA)

## Contributing
Feel free to fork the repository and submit pull requests.

## License
This project is licensed under the MIT License.

