# Badgerloop Software: Dashboard

*Contributors: Vaughn Kottler, Chase Schachenman, Gregg Van Dycke*

More info coming soon.

Pod I Dashboard: [https://github.com/vkottler/software](https://github.com/vkottler/software) 

Set up Go: https://golang.org/doc/install

Download current Go Version at https://golang.org/dl/

Run to put Go files in local directory:

tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz ( replace $VERSION and $OS with what you installed)

Alow yourself to access the go comand system wide

Open /etc/profile and add the line:
export PATH=$PATH:/usr/local/go/bin

In the dashboard directory with main.go run:

go get
go build
./dashboard

Open a new terminal tab and go into ui directory and run:

npm install
nodejs server (should expose server on localhost:8080)
