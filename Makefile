default: fmt compile run

fmt:
	go fmt *.go

compile:
	go build .

run:
	-./ptz-status

qAll:
	-./ptz-status cam1
	-./ptz-status cam2
	-./ptz-status cam3
	-./ptz-status cam4
	-./ptz-status cam5
