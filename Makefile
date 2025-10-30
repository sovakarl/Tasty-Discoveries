SourceDir=src
RunAppDir=${SourceDir}/cmd



default: 
	go run ${RunAppDir}/main.go

compose:
	docker compose up -d 





