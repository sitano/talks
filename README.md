This repository holds various talks that may be viewed with the present tool.

To install the present tool, use go get:

	go get golang.org/x/tools/cmd/present

To deploy these talks to App Engine, copy the tools/cmd/present directory to
the root of this repository and run:

	appcfg.py update -A your-app-id -V your-app-version /path/to/talks
