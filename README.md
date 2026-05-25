# Rate Limiter - Token Bucket

A lightweight and efficient Token Bucket rate limiter implementation.

This project provides a simple way to protect APIs and services from abuse by limiting the number of requests allowed within a specific time frame.

---

## Features

* Token Bucket rate limiting algorithm
* Lightweight and fast
* Easy integration
* Concurrent-safe implementation
* Configurable bucket capacity and refill rate
* Suitable for APIs, authentication systems, and backend services

---

## What is the Token Bucket Algorithm?

The Token Bucket algorithm is a commonly used rate limiting technique.

* A bucket contains a fixed number of tokens.
* Tokens are added back into the bucket at a steady refill rate.
* Every request consumes a token.
* If no tokens remain, the request is rejected or delayed.

This approach allows controlled bursts of traffic while maintaining a stable average request rate.

---

## Use Cases

* API rate limiting
* Preventing brute-force attacks
* Limiting login attempts
* Protecting backend resources
* Controlling traffic spikes
* Distributed service protection

---

## Contributing

Contributions are welcome.

Feel free to fork the repository, open issues, or submit pull requests.

---

## License

MIT License

---
