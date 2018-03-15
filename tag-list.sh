#! /usr/bin/env sh

help_text='docker run -it tag-list REPO'

if [[ ! $1 ]]; then
  echo 'Please enter a repository name. Example: '\''hello-world'\''.'
  exit 1
fi

# Set tags URL
base_url='https://registry.hub.docker.com/v2/repositories'
if [[ "$1" =~ '/' ]]; then
  repo_name="$1"
else
  repo_name="library/$1"
fi
full_url="$base_url/$repo_name/tags/"


tags="$(curl -s -S "$full_url")"

if [[ "$tags" =~ "Object not found" ]]; then
  echo 'Repo '\'"$1"\'' not found'
  exit 1
fi

if [[ "$2" == '--debug' ]]; then
  echo "$tags"
else
  echo "$tags" | jq --raw-output '["TAG","ARCHITECTURE","OS"], (."results"[] | [."name", ."images"[]."architecture", ."images"[]."os"]) | @tsv' | column -t
fi
