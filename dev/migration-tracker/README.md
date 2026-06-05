# KCC Direct Migration Tracker

This folder contains a static website and metadata file for tracking the migration of KCC resources from Terraform/DCL to Direct Controllers.

## Files

- `data.json`: The metadata file containing all the KCC resources and their migration status, dependencies, and steps.
- `index.html`: The static webpage to display the data.
- `app.js` & `style.css`: The UI logic and styles.
- `generate_data.py`: A Python script to scaffold `data.json` initially from `hack/resource-dependencies.md`.

## How to View Locally

You can run any simple HTTP server to serve the static files. For example, using Python 3:

```bash
cd dev/migration-tracker
python3 -m http.server 8000
```

Then open [http://localhost:8000](http://localhost:8000) in your web browser.

## Updating Data

To update the status of a migration, you can edit `data.json` manually (e.g. for notes, mock status, etc.).

However, to automatically detect migration progress (generated types, reference files, direct controller implementations, test fixtures) and keep dependencies, sort orders, and resource lists up to date:

```bash
cd dev/migration-tracker
python3 generate_data.py
```

This script reads the existing `data.json` and merges manual updates (like notes or mock progress) with the auto-detected progress from the codebase.

