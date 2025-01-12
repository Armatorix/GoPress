backend-merge:
	-rm all.yaml
	find ./*.yaml -type f -exec swagger-cli validate {} \;
	yq ea '. as $$item ireduce ({}; . * $$item )' ./*.yaml > all.yaml
	swagger-cli validate all.yaml
