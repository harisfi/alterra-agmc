testcov:
	rm profile.cov && clear && go test -v -coverpkg=./... -coverprofile=profile.cov ./...

showcov:
	clear && go tool cover -func profile.cov