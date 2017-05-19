package = github.com/Farmerbank/skillserver

.PHONY: install release test

install:
	go get -t -v ./...

release:
	mkdir -p release
	rm -rf release/*
	GOOS=linux GOARCH=amd64 go build -o release/skillserver-linux-amd64 $(package)
	# GOOS=linux GOARCH=386 go build -o release/skillserver-linux-386 $(package)
	# GOOS=windows GOARCH=amd64 go build -o release/skillserver-windows-amd64.exe $(package)
	# GOOS=windows GOARCH=386 go build -o release/skillserver-windows-386.exe $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/skillserver-osx-amd64 $(package)
test:
	go test -v


run:
	go build ./...
	go install ./...

# clean:
# 	rm -rf testfiles
# 	[ -f ./testfile.zip ] && rm ./testfile.zip || echo testfile.zip does not exist
# 	[ -f ./barney_report.json ] && rm ./barney_report.json || echo barney_report.json does not exist
# 	rm -rf *.json
# 	rm -rf *.sig


# runSign: run generateKeyPair runArchScan
# 	barney --verbose sign ./comparearch.json
# 	barney --verbose verify ./comparearch.json

# runScanDir:	run
# 	barney --verbose dir ./testfiles --outfile ./comparearch.json

# runArchScan: run generateTestFilesZip
# 	barney --verbose arch ./testfile.zip --outfile ./comparearch.json

# runArchScanlarge: run generateTestFilesZipLarge
# 	barney --verbose arch ./testfile.zip

# runCompare: runScanDir
# 	barney --verbose compare --archive ./testfile.zip ./barney_report.json --outfile ./comparearch.json
# 	barney --verbose compare --directory ./testfiles ./barney_report.json --outfile ./comparedir.json

# fullRunWithSignVerify: run generateKeyPair generateTestFilesZip
# 	barney --verbose arch ./testfile.zip --sign --private_key private_key.pem --outfile ./signed_compare_arch.json --signature_file sig_comp_arch.sig
# 	barney --verbose dir ./testfiles --sign --private_key private_key.pem --outfile ./signed_compare_dir.json --signature_file sig_comp_dir.sig
# 	barney --verbose compare ./signed_compare_arch.json --archive ./testfile.zip --outfile barney_arch_report.json --public_key public_key.pem --signature_file sig_comp_arch.sig
# 	barney --verbose compare ./signed_compare_dir.json --directory ./testfiles --outfile barney_arch_report.json --public_key public_key.pem --signature_file sig_comp_arch.sig



# generateKeyPair:
# 	[ -f ./private_key.pem ] && rm ./private_key.pem || echo private_key.pem does not exist
# 	[ -f ./public_key.pem ] && rm ./public_key.pem || echo public_key.pem does not exist
# 	openssl genrsa -out ./private_key.pem 2048
# 	openssl rsa -pubout -in ./private_key.pem -out ./public_key.pem

# generateTestFiles: clean
# 	mkdir ./testfiles
# 	mkdir ./testfiles/one
# 	mkdir ./testfiles/three
# 	mkdir ./testfiles/two
# 	for i in $$(seq 1 20000); do dd if=/dev/urandom bs=1 count=1 of=./testfiles/testfile$${i} ; done
# 	for i in $$(seq 1 20000); do dd if=/dev/urandom bs=1 count=1 of=./testfiles/one/testfile$${i} ; done
# 	for i in $$(seq 1 10); do dd if=/dev/urandom bs=512 count=5000 of=./testfiles/two/testfile$${i} ; done
# 	for i in $$(seq 1 100); do dd if=/dev/urandom bs=512 count=5000 of=./testfiles/three/testfile$${i} ; done

# generateShortTestFiles: clean
# 	mkdir ./testfiles
# 	mkdir ./testfiles/one
# 	for i in $$(seq 1 200); do dd if=/dev/urandom bs=1 count=1 of=./testfiles/testfile$${i} ; done
# 	for i in $$(seq 1 200); do dd if=/dev/urandom bs=1 count=1 of=./testfiles/one/testfile$${i} ; done

# generateTestFilesZip: generateShortTestFiles
# 	zip -r testfile.zip testfiles

# generateTestFilesZipLarge: generateTestFiles
# 	zip -r testfile.zip testfiles


# generateLargeTestFiles:
# 	mkdir -r ./testfiles/large

# 	for i in $$(seq 1 100); do dd if=/dev/urandom bs=512 count=50000 of=./testfiles/$${x}/testfile$${i} ; done\