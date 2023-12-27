run:
	go run .

build:
	mkdir -p builds
	go build -o builds/backendboilerplate .

install:
	@echo "Still working on this"
	cp builds/backendboilerplate/* /usr/bin/
	backendboilerplate init

debs:
	mkdir -p builds
	go build -o builds/backendboilerplate .
	mkdir -p installer/backendboilerplate_1.2-1_amd64/usr/local/bin
	cp builds/backendboilerplate installer/backendboilerplate_1.2-1_amd64/usr/local/bin/
	mkdir -p installer/backendboilerplate_1.2-1_amd64/DEBIAN
	cp deb/p* installer/backendboilerplate_1.2-1_amd64/DEBIAN/
	cp deb/control_amd64/* installer/backendboilerplate_1.2-1_amd64/DEBIAN/
	dpkg-deb --build --root-owner-group installer/backendboilerplate_1.2-1_amd64
	
	mkdir -p installer/backendboilerplate_1.2-1_i386/usr/local/bin
	cp builds/backendboilerplate installer/backendboilerplate_1.2-1_i386/usr/local/bin/
	mkdir -p installer/backendboilerplate_1.2-1_i386/DEBIAN
	cp deb/p* installer/backendboilerplate_1.2-1_i386/DEBIAN/
	cp deb/control_i386/* installer/backendboilerplate_1.2-1_i386/DEBIAN/
	dpkg-deb --build --root-owner-group installer/backendboilerplate_1.2-1_i386

	mkdir -p installer/backendboilerplate_1.2-1_arm64/usr/local/bin
	cp builds/backendboilerplate installer/backendboilerplate_1.2-1_arm64/usr/local/bin/
	mkdir -p installer/backendboilerplate_1.2-1_arm64/DEBIAN
	cp deb/p* installer/backendboilerplate_1.2-1_arm64/DEBIAN/
	cp deb/control_arm64/* installer/backendboilerplate_1.2-1_arm64/DEBIAN/
	dpkg-deb --build --root-owner-group installer/backendboilerplate_1.2-1_arm64