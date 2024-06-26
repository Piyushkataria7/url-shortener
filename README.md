# url-shortener

This repository contains a simple URL shortener service implemented in Go, focusing on providing a REST API for URL shortening, redirection, and domain metrics.

Features:
Shorten URL API: Accepts a long URL and generates a short URL.
Redirection API: Redirects users from a short URL to the original long URL.
In-memory Storage: URLs and their mappings are stored in memory.
Metrics API: Provides metrics on top 3 domains shortened the most.

Endpoints:
POST /shorten: Shorten a URL.
GET /topdomains: Get top domains by number of shortenings.
GET /{shortURL}: Redirects to original URL.


Example
Shorten URL

Request: curl -X POST -H "Content-Type: application/json" -d '{"url": "https://example.com"}' http://localhost:8080/shorten

Response: {"short_url":"http://localhost:8080/abc123"}

Redirect
Request: curl -i http://localhost:8080/abc123

Response: Location: https://example.com

Top Domains
Request: curl http://localhost:8080/topdomains

Response:[
  "example.com: 1"
]

Dependencies
Go 1.16 or higher

Standard library packages (no external dependencies)