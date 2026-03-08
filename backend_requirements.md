# Backend Requirements Specification: URL Shortener

This document outlines the logical architecture, functional requirements, and non-functional performance standards required for the backend infrastructure of the scalable URL shortener platform.

## 1. Functional Requirements (System Capabilities)
**Goal:** Define the exact operations the backend API and background workers must perform.

* **Link Generation & Management:**
  * Accept a long URL payload and return a unique, shortened alias.
  * Validate that incoming long URLs are well-formed and safe (e.g., block loopback addresses or known malicious domains).
  * Support custom, user-defined short aliases, ensuring they do not conflict with existing links.
  * Support optional expiration timestamps; links past this time must return a 404 or a custom expired page.
* **Redirection Engine:**
  * Intercept requests for a short code, locate the corresponding long URL, and execute an HTTP 301 (Permanent) or 302 (Temporary) redirect.
* **Telemetry & Analytics Ingestion:**
  * Asynchronously record metadata for every successful redirect without delaying the user's request.
  * Extract and store data points including: Timestamp, IP address (for geospatial translation), User-Agent (browser/device), and HTTP Referrer.
* **Authentication & Authorization:**
  * Provide secure endpoints for user registration, login, and session management.
  * Enforce role-based access control (users can only edit/view analytics for their own links).



## 2. Non-Functional Requirements (System Performance)
**Goal:** Define the strict constraints the system must operate within under heavy load.

* **Latency:**
  * The Read Path (redirection) must be highly optimized. The backend must process the redirect in under 10 milliseconds (excluding network transit time).
* **High Availability & Fault Tolerance:**
  * The redirection service must be resilient. If an application server crashes, the system must continue routing traffic to healthy nodes.
* **Scalability:**
  * The system must be capable of handling an extreme read-to-write ratio (e.g., 100:1).
  * The system must survive "Hot Key" events (viral traffic spikes hitting a single short link simultaneously) without exhausting database connections.
* **Data Integrity:**
  * The system must mathematically or structurally guarantee that a short code is never assigned to two different long URLs (Zero Collisions).
* **Rate Limiting:**
  * Implement strict rate limits on the Write Path (link creation) to prevent automated bot abuse or database exhaustion.

## 3. Logical Architecture Components
**Goal:** Define the necessary subsystems and their responsibilities, regardless of the underlying technology stack.

* **The API Gateway / Routing Layer:**
  * Acts as the single entry point. Handles SSL termination, rate limiting, and routes requests to the appropriate internal services.
* **The Application Service (Stateless Handlers):**
  * Contains the core business logic. Validates payloads, queries the cache/database, and formats HTTP responses. Must be completely stateless to allow horizontal scaling.
* **The Key Generation Service (KGS):**
  * An isolated background process responsible for pre-computing millions of unique, randomized short codes. It ensures the application servers never have to "guess" an available code or resolve collisions during a live user request.
* **The Caching Layer (In-Memory Storage):**
  * Stores the most frequently accessed `short_code -> long_url` mappings. Must be configured with an eviction policy (e.g., Least Recently Used) to manage limited RAM capacity.
* **The Persistent Data Store:**
  * The definitive source of truth. Stores user accounts, link mappings, and the raw analytics data. Must enforce strict uniqueness constraints on the short codes.
* **The Asynchronous Event Queue (Optional but Recommended):**
  * A messaging broker that sits between the redirection engine and the database. It catches the analytics data (clicks) and queues it up to be written to the database in batches, ensuring the database isn't overwhelmed by concurrent write requests during a viral event.
