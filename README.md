# 🌤️ Weather App

A mobile weather application built with **Flutter** for the frontend and **Go** for the backend. It fetches real-time weather data and manages user interactions with a robust and containerized backend.

---

![Screenshot 2025-05-06 174428](https://github.com/user-attachments/assets/b7a40efc-308f-4238-8867-6fbf620e143d)

![Screenshot 2025-05-06 174436](https://github.com/user-attachments/assets/c5b1dddd-7149-428c-8c3d-f40f1110ad84)

---

## 🚀 Features

- Real-time weather data using OpenWeather API
- Clean and responsive Flutter UI
- Go backend for efficient API handling
- Email support (via Mailtrap)
- Docker Compose for easy backend setup

---

## 🧱 Project Structure

/
├── docker-compose.yml
├── src/
│ ├── client/ # Flutter frontend
│ └── server/ # Go backend

---

## 🔧 Prerequisites

- [Flutter](https://flutter.dev/docs/get-started/install)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Go](https://golang.org/dl/) (for local development or debugging)
- [air](https://github.com/cosmtrek/air) (live reload tool for Go)
- [Make](https://www.gnu.org/software/make/) – To simplify development workflows using predefined `Makefile` commands

---

## Credits

- CodeHQ (UI Design): https://www.youtube.com/watch?v=msoKuk-5QFg

---

## ⚙️ Backend Setup

1. **Add a `.env` file to `/src/server/` directory:**

   ```env
   export OPEN_WEATHER_API_KEY=your_openweather_api_key
   export DB_ADDR=your_db_connection_string
   export MAILTRAP_USERNAME=your_mailtrap_username
   export MAILTRAP_PASSWORD=your_mailtrap_password
   ```

2. **Start the backend services with Docker Compose:**

   From the project root directory, run:

   ```bash
    cd src/server
    docker-compose up --build
    air
   ```

3. **Migrate your database:**
   ```bash
   make migrate-up
   ```

## 📱 Frontend Setup (Flutter)

1. **Navigate to the Flutter project:**
  ```bash
    cd src/client
  ```

2. **Install dependencies**:
  ```bash
    flutter pub get
  ```

3. **Run the Flutter app**:
  ```bash
  flutter run
  ```

Make sure an emulator or a physical device is connected.
