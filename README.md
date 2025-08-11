# Daily Tasks

Daily tasks is an application, which allows to create some tasks for yourself and log the progress.

The basic functionality would just show the progress using web or mobile UI on pages like week, month, year progress.

|                | M      | T      | W      | T      | F | S | S |
|----------------|--------|--------|--------|--------|---|---|---|
| "water plants" | done   | done   | done   | -      | - | - | - |
| "read book"    | 15 p   | 32 p   | 16 p   | 25 p   | - | - | - |
| "run"          | 30 min | 35 min | 40 min | 35 min | - | - | - |
| "go to bed"    | 1:30   | 1:40   | 1:00   | 0:55   | - | - | - |

In the future, I am planning to add some statistics to the data.

## Deployment

### Using Docker Compose (Development only)

The entire application (backend, database, and frontend) can be started with a single command:

```bash
docker-compose up
```

This will:
1. Start a PostgreSQL database
2. Run db migrations if it is the first time running this container 
3. Build and start the Go backend server
4. Build and start the Expo frontend application

The services will be available at:
- Backend API: http://localhost:8080
- Metro for Web: http://localhost:8081

To stop all services:

```bash
docker-compose down
```
