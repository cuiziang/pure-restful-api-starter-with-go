# Pure RESTful API Starter with Go


## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Features

- Primarily uses Go's standard library packages for core functionality 
- Minimal external dependencies, only uses `github.com/go-sql-driver/mysql`
- Follows the standard Go project layout
- Easy to extend with additional dependencies
- Includes some intentionally redundant code, you can refactor it to suit your needs

## Requirements

- Go 1.23 or higher
- MySQL database

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/cuiziang/pure-restful-api-starter-with-go
   ```

2. Change to the project directory:
   ```
   cd pure-restful-api-starter-with-go
   ```
3. Set up environment variables for database connection:
    ```
    export DB_HOST=your_database_host
    export DB_NAME=your_database_name
    export DB_USER=your_database_user
    export DB_PASS=your_database_password
    ```
    Note: For Windows users, use set instead of export.


3. Install the dependencies, build the project, and run the server:

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT) - see the [LICENSE](LICENSE) file for details.
