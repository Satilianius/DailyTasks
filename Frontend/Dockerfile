# Frontend Dockerfile for Expo dev server
# Use Node LTS
FROM node:20-bullseye

# Set working directory
WORKDIR /app

# Install dependencies first (leverage Docker layer caching)
COPY package.json package-lock.json* ./
RUN npm ci || npm install

# Copy the rest of the application
COPY . .

# Environment for reliable file watching in containers
ENV CHOKIDAR_USEPOLLING=true \
    WATCHPACK_POLLING=true

# Expose Expo ports (web and Metro)
EXPOSE 19006 8081

# Default command: run Expo web server accessible from outside the container
# Binding to 0.0.0.0 ensures it's reachable via host port mapping
CMD ["npm", "run", "web", "--", "--host", "0.0.0.0", "--port", "19006"]
