# Frontend

React + Vite frontend for Kairo.

## Prerequisites

- Node.js 22+
- Workspace dependencies installed at repository root

## Local development

From `/home/runner/work/kairo/kairo`, install dependencies:

```bash
npm install
```

Build required workspace packages for frontend imports:

```bash
cd packages/zod && npm run build
cd ../openapi && npm run build
```

Start the frontend:

```bash
cd /home/runner/work/kairo/kairo/apps/frontend
npm run dev
```

The `dev` script waits for `packages/openapi/dist/index.js`, so the OpenAPI package must be built first.

## Scripts

Run from `/home/runner/work/kairo/kairo/apps/frontend`:

- `npm run dev`
- `npm run build`
- `npm run test`
- `npm run lint`
- `npm run lint:fix`
- `npm run format`
- `npm run format:fix`
- `npm run typecheck`
