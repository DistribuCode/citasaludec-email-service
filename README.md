# 📧 Email Service

This microservice is responsible for sending email notifications across the system — including appointment confirmations, cancellations, and general communication. It's written in **Go** using idiomatic project structure with packages and `internal` modules. The service is containerized with Docker and set up for automated deployment via GitHub Actions to an EC2 instance.

---

## 🧩 Features

- Send emails for various system events  
- Modular Go structure using packages and `internal/`  
- Middleware support and structured routing  
- Configurable via environment variables  
- Runs in a Docker container  
- CI/CD enabled via GitHub Actions  
- Deployable on AWS EC2

---

## 📁 Project Layout

- `cmd/main.go`: Entry point of the service  
- `internal/`: Contains all main application logic
  - `config/`: App configuration and env loader
  - `controllers/`: Handles request logic and email triggers
  - `middlewares/`: Middleware (e.g., auth or logging)
  - `models/`: Structs and data models
  - `routes/`: HTTP routing logic
  - `services/`: Email sending logic (SMTP, API clients, etc.)
- `.email-service.env`: Optional env file just for this service
- `Dockerfile`: Build configuration
- `go.mod / go.sum`: Go module dependencies

---

## 🔐 Environment Variables (`.env`)

```env
PORT=4003
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=youremail@example.com
SMTP_PASS=yourpassword
SENDER_EMAIL=no-reply@example.com

You can use .env or .email-service.env in development.
🐳 Docker Usage
1. Build the Docker image

docker build -t email-service .

2. Run the container

docker run -d -p 4003:4003 --env-file .env email-service

🔌 API Endpoints
Method	Route	Description
POST	/send-email	Sends an email notification

    Expected payload (JSON):

{
  "to": "user@example.com",
  "subject": "Appointment Confirmed",
  "body": "Your appointment is confirmed for tomorrow."
}

☁️ EC2 Deployment
1. Connect to your EC2 instance

ssh -i key.pem ec2-user@<YOUR_EC2_PUBLIC_IP>

2. Pull the image

docker pull jeffri1997/email-service:qa

3. Run the service

docker run -d -p 4003:4003 --env-file .env jeffri1997/email-service:qa

✅ Ensure port 4003 is open in your EC2 security group.
🤖 GitHub Actions CI/CD

A workflow is set up to automatically build and push the image on qa branch updates:

name: Build and Push Email Service

on:
  push:
    branches: [qa]

jobs:
  build-and-push:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Docker Login
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      - name: Build and Push
        run: |
          docker build -t jeffri1997/email-service:qa .
          docker push jeffri1997/email-service:qa

🧪 Testing

If you include tests in a tests/ folder, run them with:

go test ./...

You can also test individual handlers with tools like GoConvey or Testify.
📬 Example Email Service Use

    Appointment booked → send confirmation

    Appointment cancelled → send cancellation notice

    System notifications or alerts

Integrates well with microservices using HTTP or message queues.
👤 Author

Jefferson Marcalla
GitHub: @Jeff97ares