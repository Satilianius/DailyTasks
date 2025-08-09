FROM node:24-alpine AS builder

WORKDIR /app

# Install dependencies first (better layer caching)
COPY package*.json ./
RUN npm ci

# Copy source
COPY . .

# Expose Expo/Metro and Web ports
EXPOSE 19000
EXPOSE 19001
EXPOSE 19002
EXPOSE 19006
EXPOSE 8081

# Run Expo Web, binding to all interfaces for access from host
CMD ["npm", "run", "web"]
