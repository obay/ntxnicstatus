# Change Nutanix VM NIC Connectivity Status

This is a simple command to facilitate changing the NIC connectivity of a given virtual machine in a Nutanix cluster using the Prism Elements API (V2).

While you can do that directly from the Prism UI, having a command line allows you to automate the process in your monitoring software and trigger it whenever needed.

## Use Case

This command can be helpful if you have a virtual machine with questionable reliability and you want to have another identical VM ready to take over. Think of as a kind of poor man fail-over.

You can simply clone the VM, keep the clone running with a disconnected NIC. Monitor the master VM and once it becomes unresponsive, use this command to disconnect the network card and connect the clone VM.

While Nutanix will cover any case where the host/node fails, Nutanix doesn't manage the application layer and hence that is the responsibility of the application administrator. This command-line gives application administrator extended reach to control the infrastructure which can be useful if you are running an application that doesn't have failover functionality built-in at the application layer.

## Fun Fact

Running this command on your LAN should take somewhere between 0.9 to 1.7 seconds to complete. If your monitoring software is responsive enough, your users shouldn't be down for more than a few seconds.

## Installation

### Windows

1. Download https://github.com/obay/ntxnicstatus/blob/master/bin/ntxnicstatus-0.1-windows-amd64.zip
2. Unzip ntxnicstatus-0.1-windows-amd64.zip
3. Move ntxnicstatus to C:\Windows\system32
4. Click on Start --> Run...
5. Type cmd and hit Enter
6. Run `ntxnicstatus --version`. You should see
   ntxnicsstatus version 0.1

### MacOS

1. Download https://github.com/obay/ntxnicstatus/blob/master/bin/ntxnicstatus-0.1-darwin-amd64.zip
2. Unzip ntxnicstatus-0.1-darwin-amd64.zip
3. Move ntxnicstatus to /usr/loca/bin by running the following command in the terminal:
   `sudo mv ~/Downloads/ntxnicstatus /usr/loca/bin`
4. Run `ntxnicstatus --version`. You should see
   ntxnicsstatus version 0.1

### Linux

1. Download https://github.com/obay/ntxnicstatus/blob/master/bin/ntxnicstatus-0.1-linux-amd64.zip
   `wget https://github.com/obay/ntxnicstatus/blob/master/bin/ntxnicstatus-0.1-linux-amd64.zip`
2. Unzip ntxnicstatus-0.1-linux-amd64.zip
   `unzip ntxnicstatus-0.1-linux-amd64.zip`
3. Move ntxnicstatus to /usr/loca/bin by running the following command in the terminal:
   `sudo mv ntxnicstatus /usr/loca/bin`
4. Run `ntxnicstatus --version`. You should see
   ntxnicsstatus version 0.1

## Usage

The following parameters are used when running the command

### --hostname

Required. This is the hostname, IP address, or fully qualified domain name (FQDN) for any of your Nutanix CVMs or the virtual IP address of the entire Nutanix cluster. It is the same IP address you use to access Prism.

### --port

This is optional. The default used is the default Nutanix Prism port number which is 9440. You don't have to provide this port number if you are not using a special configuration that differs from the usual setup.

### --username

Required. This is the username you use to login to Prism Elements. It is a good practice to create a username especially for use with scripts and automated tasks to understand who did what when things go wrong (which they eventually tend to do)

### --password

Required. Password used to access Prism Elements. It is a good practice to use a long (more than 12) characters, randomly generated password which contains upper characters, lower characters, digits, and special characters just to be safe. Use a password database to keep all your passwords safe and don't reuse passwords if you don't feel like being hacked.

### --insecure

Optional. Default is set to false which means secure. This defines if your Nutanix Prism Elements is using a self-signed (insecure) SSL certificate or you do have a proper SSL certificate (which you should do to improve security).

### --vmname

Required. Name of the virtual machine that you want to connect or disconnect from the network. This value is case sensitive. Make sure there are no multiple virtual machines with the same name. Multiple virtual machines with the same name is not a case that is handled at this version.

### --mac

Required if your virtual machine has more than 1 NIC attached to it. Not required if you have 1 NIC only attached to the virtual machine. Format of the MAC address can be either ':' separated or '-' separated. The example below is ':' separated

## --connected

Optional. Default is true. This is the status of the NIC. If you want to have the NIC connected, set this to true. If you want to disconnect the NIC, set this to false.



## Example

```bash
ntxnicstatus --hostname='192.168.20.20' --port=9440 --username='admin' --password='supersecret' --insecure=false --vmname='Windows VM test' --mac='50:6b:8d:57:a3:81' --connected=false
```



## Disclaimer

This code is provided as-is for anyone using Nutanix. This is not a commercially supported software and is not associated with Nutanix in any way. Use at your own risk.
