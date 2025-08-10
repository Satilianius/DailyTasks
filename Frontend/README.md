# DailyTasks Frontend (Expo + Expo Router)

This is the frontend app for DailyTasks built with Expo and Expo Router. It supports running in the browser (web) as well as on Android and iOS via Expo.

## Prerequisites
- Node.js 18 or 20 (LTS recommended)
- npm 9+ (comes with Node) or another package manager
- Optional (for devices/emulators):
  - Android Studio / Emulator, or a physical Android device with Expo Go
  - Xcode / iOS Simulator on macOS, or a physical iOS device with Expo Go

## Install dependencies
```bash
npm install
```

## Run the app
You can run the app in different modes using the provided npm scripts.

- Start (choose a platform in Expo Dev Tools):
```bash
npm start
```

- Web (browser):
```bash
npm run web
```
  Access: http://localhost:8081/

- Android (launches Expo on an Android emulator or Expo Go on a connected device):
```bash
npm run android
```

- iOS (launches Expo on iOS Simulator or Expo Go on a connected device; macOS only):
```bash
npm run ios
```

Expo Dev Tools will open in your browser. For mobile devices, you can scan the QR code with Expo Go.

## Tests
Run unit tests with Jest:
```bash
npm test
```

## Docker (web dev server)
A Dockerfile is provided to run the web dev server inside a container.

1) Build the image:
```bash
docker build -t dailytasks-frontend .
```

2) Run the container (bind ports for web and Metro):
```bash
docker run --rm -p 8081:8081 -p 8081:8081 dailytasks-frontend
```

Then open http://localhost:8081/ in your browser.

Note: The container command runs `npm run web` and binds to 0.0.0.0 for external access.

### Docker Compose (from repository root)
If you're using the provided docker-compose.yml at the repo root, you can start only the frontend service:
```bash
docker compose up frontend
# or with legacy syntax
docker-compose up frontend
```
Then open http://localhost:8081/ in your browser.

## Environment / Backend configuration
This project currently has no explicit API base URL configuration in the codebase. If you need to connect to a backend API:
- Add your environment handling (e.g., via a config file or using `process.env` with `app.config`/`app.json`).
- Document the required variables (e.g., `API_BASE_URL`) and how to set them.

## Troubleshooting
- If the Metro bundler or web server misbehaves, clear caches:
```bash
npx expo start -c
```
- Ensure no other process is using port 8081 (Metro).
- For Android/iOS, verify that emulators are running or devices are connected and recognized by `adb devices` / Xcode.

## Project scripts
// language=json
```json
{
  "scripts": {
    "start": "expo start",
    "android": "expo start --android",
    "ios": "expo start --ios",
    "web": "expo start --web",
    "test": "jest --watchAll"
  }
}
```
