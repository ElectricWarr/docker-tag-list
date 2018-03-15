. project.cfg

full_repo="$dockerhub_namespace/$repo"

# Tag
docker tag "$repo:$tag" "$full_repo:$tag"
docker tag "$repo:$tag" "$full_repo:latest"

# Interactive Login
docker login || exit 1

# Push
docker push "$full_repo:$tag"
docker push "$full_repo:latest"
