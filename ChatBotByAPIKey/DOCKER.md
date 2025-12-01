# AI ChatBot Docker Setup

## Build Docker Image

```bash
docker build -t aichatbot:latest .
```

## Run Container

With `.env` file (baked into image):
```bash
docker run -it aichatbot:latest
```

With environment variable override:
```bash
docker run -it -e GENAI_API_KEY=your_api_key_here aichatbot:latest
```

With mounted .env file:
```bash
docker run -it --env-file .env aichatbot:latest
```

## Interactive Mode

The container runs in interactive mode (`-it`) allowing you to send messages to the ChatBot from stdin.

## Notes

- The Dockerfile uses a multi-stage build to minimize final image size
- Alpine Linux is used as the base for a lightweight image
- Set the `GENAI_API_KEY` environment variable with your Google Generative AI API key
- The application writes files to the container's `/root/Desktop` directory
