#!/bin/bash
: <<'COMMENT'
This script starts the Labspace environment using Docker Compose. It sets the CONTENT_PATH
environment variable to the current working directory and uses a specific Labspace
COMMENT

CONTENT_PATH=$PWD docker compose -f oci://dockersamples/labspace-content-dev -f .labspace/compose.override.yaml up