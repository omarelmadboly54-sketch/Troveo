# Trovero V1 frontend

A responsive React + Vite storefront/authentication/admin frontend for the supplied Go/Gin backend.

## Run locally

```bash
cd trovero_frontend
cp .env.example .env
npm install
npm run dev
```

In local development, Vite proxies `/api` to `http://localhost:8080`. Set `VITE_API_BASE_URL` to your deployed backend API root (including `/api`) for production. The supplied Gin backend does not currently configure CORS, so production deployments should use a same-origin reverse proxy or add a restricted CORS policy.

## Backend contract verified from the supplied source

| Method | Path | Access | Request / response |
|---|---|---|---|
| POST | `/api/users` | Public | `{ first_name, last_name, email, password, account_type }`; account type must be `seller` or `customer`; password min 8; success `201 { message }` |
| POST | `/api/login` | Public | `{ email, password }`; success `200 { accessToken }` |
| GET | `/api/users` | Bearer JWT + role `admin` | Query: `page` (default 1), `limit` (default 10), optional `account_type`; response `200 { users: [...] }` |
| GET | `/api/users/:id` | Bearer JWT + role `admin` | Response `200 { user: {...} }` |
| DELETE | `/api/users/:id` | Bearer JWT + role `admin` | Success `200 { message }`; backend only deletes rows with `role='user'` |

User response fields are `id`, `first_name`, `last_name`, `email`, `role`, and `account_type`. Password and timestamps are not exposed. Admin list filtering uses `account_type=seller` or `account_type=customer` on the same `/users` endpoint.

## Important backend limitations (not fabricated in this frontend)

- There are no product, cart, product-settings, or general-settings endpoints in the supplied backend. Product and settings destinations therefore show an explanatory placeholder.
- The list response contains only `{ users: [...] }`, with no total count, total pages, or pagination metadata. The UI uses `page` and `limit=10`; it enables Next when the returned page contains 10 users and cannot display an exact total-page count.
- The only public registration account types are `seller` and `customer`. Admin accounts are not creatable through the public registration endpoint.
- The JWT payload includes `user_id`, `role`, `issued_at`, and `expired_at`; the frontend decodes the role for navigation. The backend remains the authority and protects all admin APIs with JWT role middleware.
- The supplied project includes a local development database connection and signing secret in `Backend/app.env`. Do not expose or reuse those values in production; rotate the signing secret and use environment-managed credentials.

## Branding

`public/trovero-logo.png` is derived from the provided logo artwork with the neutral checkerboard removed to create a transparent PNG.
