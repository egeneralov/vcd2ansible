~/go/bin/vcd2ansible:
	go build -v -o ~/go/bin/vcd2ansible cmd/vcd2ansible/*.go
clean:
	rm -f ~/go/bin/vcd2ansible