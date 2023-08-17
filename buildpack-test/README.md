# Building

`pack build hello --buildpack paketo-buildpacks/go --env BP_GO_BUILD_LDFLAGS="-X main.who=buildpack" --builder paketobuildpacks/builder:base`
