#!/bin/sh

set -x

# Cleanup
rm -rf docker-registry

# Setup the environment
export SETTINGS_FLAVOR=test
export DOCKER_REGISTRY_CONFIG=config_test.yml

# Get latest docker registry
git clone -q https://github.com/dotcloud/docker-registry.git
cd docker-registry

# Get dependencies
pip install -q -r requirements.txt
pip install -q -r test-requirements.txt
pip install -q tox

# Run registry tests
tox || exit 1
export PYTHONPATH=$(pwd)/docker-registry
python -m unittest discover -p s3.py -s test || exit 1
python -m unittest discover -p workflow.py -s test

