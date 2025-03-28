# Centralized Finance Dashboard

## Setup

This guide will help you clone the repository, install Docker, and run the application using Docker Compose.

### Prerequisites

- **Git:**  
  Make sure Git is installed on your machine. You can download it from [git-scm.com](https://git-scm.com/downloads).

- **Docker:**  
  Install Docker from [Docker's official website](https://www.docker.com/get-started).  
  *Note:* Docker Desktop includes Docker Compose, so no separate installation is needed.

### Clone the Repository

1. Open your terminal.
2. Run the following command to clone the repository:
   ```bash
   git clone https://github.com/yourusername/your-repo-name.git
   ```
3. Navigate into the repository directory:
    ```bash
    cd your-repo-name
    ```

### Running the Application with Docker Compose

Our project uses a monorepo structure with a frontend (Next.js) and a backend (Go) service.

1. Ensure you're in the repository root (where the docker-compose.yml file is located).

2. Build and start the containers with:
    ```bash
    docker-compose up --build
    ```
Docker Compose will build both the frontend and backend images and start the containers.

### Access the Application

- Frontend (Next.js):
Open your browser and go to http://localhost:3000.

- Backend (Go):
Open your browser or API client and navigate to http://localhost:8080
or access a specific endpoint like http://localhost:8080/api/data (if applicable).