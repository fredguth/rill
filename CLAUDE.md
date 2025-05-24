# Claude Code Initialization

This file contains commands and instructions that Claude Code should execute when initialized in this repository.

## Lint and Type Check Commands

- Run lint check: `npm run lint`
- Run type check: `npm run typecheck`


## Project Overview

Rill is an open source BI-as-code tool by RillData. It empowers users to build and explore dashboards with a 'BI-as-code' approach.
Key characteristics and components include:

-   **BI-as-code Core**: Data models, metrics, and dashboards are defined declaratively in YAML files within a project. This enables version control, collaboration, and reproducibility.
-   **Go Backend & Runtime**: The `runtime` (written in Go) is central to Rill. Triggered by commands like `rill start`, it parses project definitions, ingests and transforms data (often using an embedded DuckDB instance for fast local analytics), and serves data to the frontend.
-   **SvelteKit Frontend**: The user interface is built with SvelteKit and is modular:
    -   `web-local`: The primary UI for local development, offering an interactive environment for dashboard creation and data exploration.
    -   `web-common`: A library of shared UI components and utilities.
    -   `web-admin`: Provides interfaces for administrative functions, potentially for managing users or services in deployed/cloud contexts.
-   **Development vs. Production Modes**: Rill supports a development mode (via `rill start`) with full editing capabilities and a production/viewer mode for read-only dashboard access, as detailed in the "Deployment Context".
-   **Self-Hosting**: Dashboards can be self-hosted in production viewer mode, independent of Rill Cloud.


## Deployment Context

The project supports self-hosting dashboards in production viewer mode:
- Use `rill.yaml` with `prod: mode: "viewer"` configuration
- Deploy with `rill start --environment prod --readonly --no-open` for production mode
- Can be containerized with Docker and deployed to platforms like fly.io
- Production mode removes edit/deploy buttons for read-only dashboard viewing

## Key Directories

- `/admin`: Admin server and management functionality
- `/cli`: Command-line interface
- `/runtime`: Core runtime engine
- `/web-admin`, `/web-common`, `/web-local`: Web UI components
- `/docs`: Documentation

Claude Code should automatically run lint and type checks after making code changes to ensure code quality.