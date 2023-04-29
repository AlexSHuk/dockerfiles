# Build stage
#FROM python:3.9.5-slim-buster as build
#WORKDIR /app
#COPY server.py .
#RUN mkdir -p /root/.local && \
#    pip install --user --no-cache-dir -r requirements.txt && \
#    chmod -R 755 /root/.local && \
#    rm -rf ~/.cache/pip
# Run stage
#FROM python:3.9.5-slim-buster
#WORKDIR /app
#COPY --from=build /root/.local /usr/local
#COPY --from=build /app .
#CMD ["python", "server.py"]
# Build stage
FROM alpine:3.13 as build
WORKDIR /app
COPY server.py .
RUN mkdir -p /app/.local && \
    chmod -R 755 /app/.local && \
    rm -rf ~/.cache/pip

# Run stage
FROM alpine:3.13
WORKDIR /app
RUN apk update && apk add python3
COPY --from=build /app/.local /usr/local
COPY --from=build /app .
CMD ["python3", "server.py"]

