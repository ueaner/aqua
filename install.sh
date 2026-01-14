sha=""
if git diff --quiet; then
  sha=$(git rev-parse HEAD)
fi
ver=$(git tag --sort=-v:refname | head -n 1)

go install \
  -ldflags "-X main.version=$ver-local -X main.commit=$sha -X main.date=$(date +"%Y-%m-%dT%H:%M:%SZ%:z" | tr -d '+')" \
  ./cmd/aqua
