# Message Sniffer

Message sniffer is a command-line tool for displaying messages from RabbitMQ queues and exchanges.


## Getting Started


### Installation

#### From release

1. Download the appropriate binary from the list of [releases](https://github.com/joshmalbrecht/message-sniffer/releases).

2. Rename the binary how ever you would like and install the binary in a directory that exists on your $PATH (Example: `/usr/local/bin/`)

3. You should now be able to use the `message-sniffer` command

#### From source code

1. Ensure that you have your $GOPATH properly set and that the $GOPATH/bin is include in your $PATH

2. Clone the repository using `git clone https://github.com/joshmalbrecht/message-sniffer.git`

3. Run `go install` from the root of your directory

4. You should now be able to use the `message-sniffer` command

## License

[GNU General Public License v3](https://github.com/joshmalbrecht/message-sniffer/blob/main/LICENSE)