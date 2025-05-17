# golang-rest-api-template
Template for REST API made with Golang using Gin framework, PostgreSQL database, JWT authentication, Redis cache.

## Overview

This repository provides a template for building a REST API using Go with features like JWT Authentication, database operations using GORM. The application uses the Gin Gonic web framework.

## Features

- RESTful API endpoints for CRUD operations.
- JWT Authentication.
- PostgreSQL database integration using GORM.
- Redis cache.

### Authentication

To use authenticated routes, you must include the `Authorization` header with the JWT token.

```bash
curl -H "Authorization: Bearer <YOUR_TOKEN>" http://localhost:8081/api/user/1