#!/bin/bash

if ! [ -x "$(command -v cfn-lint)" ]; then
  echo '`cfn-lint` is not installed. Attempting a Brew install.' >&2
  brew install cfn-lint
fi

cfn-lint -t ./templates/*.yml && echo "Linting OK."
