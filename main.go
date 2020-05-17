// +build windows,amd64 linux,amd64 darwin,amd64

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type prismLoginDetails struct {
	username string
	password string
	hostname string
	port     int
	insecure bool
}

func prismPutRequest(prismLogin *prismLoginDetails, endpoint string, method string, data vmNicPutRequest) ([]byte, error) {
	endpoint = "https://" + prismLogin.hostname + ":" + strconv.Itoa(prismLogin.port) + "/PrismGateway/services/rest/v2.0/" + endpoint

	var req *http.Request
	// var err error
	// var json []byte

	json, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err = http.NewRequest(method, endpoint, bytes.NewBuffer(json))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(prismLogin.username, prismLogin.password)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func prismGetRequest(prismLogin *prismLoginDetails, endpoint string) ([]byte, error) {
	endpoint = "https://" + prismLogin.hostname + ":" + strconv.Itoa(prismLogin.port) + "/PrismGateway/services/rest/v2.0/" + endpoint
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(prismLogin.username, prismLogin.password)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getAllVMs(prismLogin *prismLoginDetails) (vms, error) {
	body, err := prismGetRequest(prismLogin, "vms")
	if err != nil {
		return vms{}, err
	}
	var allVMs vms
	if err := json.Unmarshal(body, &allVMs); err != nil {
		return vms{}, err
	}
	return allVMs, nil
}

func getVMNICs(prismLogin *prismLoginDetails, vmUUID string) (vmNICs, error) {
	body, err := prismGetRequest(prismLogin, "vms/"+vmUUID+"/nics")
	if err != nil {
		return vmNICs{}, err
	}
	var nics vmNICs
	if err := json.Unmarshal(body, &nics); err != nil {
		return vmNICs{}, err
	}
	return nics, nil
}

func getVMUUID(allVMs *vms, vmName string) string {
	for i := 0; i < allVMs.Metadata.Count; i++ {
		if allVMs.Entities[i].Name == vmName {
			return allVMs.Entities[i].UUID
		}
	}
	return ""
}

func getVMNetworkUUID(nics *vmNICs, nicMAC string) string {
	for i := 0; i < len(nics.Entities); i++ {
		if nics.Entities[i].MacAddress == nicMAC {
			return nics.Entities[i].NetworkUUID
		}
	}
	return ""
}

func getvmNICMAC(nics *vmNICs) (string, error) {
	if len(nics.Entities) == 1 {
		return nics.Entities[0].MacAddress, nil
	} else if len(nics.Entities) > 1 {
		return "", errors.New("Multiple NICs found. Re-run the command and define the MAC address of the network card you want to disconnect")
	}
	return "", errors.New("No NICs found for this virtual machine")
}

func setNicConnectivity(prismLogin *prismLoginDetails, vmUUID string, nicMac string, networkUUID string, connected bool) error {
	var requestBody vmNicPutRequest
	requestBody.NicID = strings.Replace(nicMac, ":", "-", 10)
	requestBody.NicSpec.IsConnected = connected
	requestBody.NicSpec.NetworkUUID = networkUUID
	jsonResponse, err := prismPutRequest(prismLogin, "/vms/"+vmUUID+"/nics/"+nicMac, "PUT", requestBody)
	if err != nil {
		return err
	}
	var xx taskResponse
	if err = json.Unmarshal(jsonResponse, &xx); err != nil {
		return errors.New(string(jsonResponse))
	}
	return nil
}

func declareFlags() {
	pflag.String("hostname", "", "Nutanix Prism hostname, IP, or FQDN (required)")
	pflag.String("username", "", "Nutanix Prism username.")
	pflag.String("password", "", "Nutanix Prism password")
	pflag.Bool("insecure", false, "Allow self signed SSL certificates. Default \"false\"")
	pflag.Int("port", 9440, "Nutanix Prism port number. Default \"9440\"")

	pflag.String("vmname", "", "Virtual machine name.")
	pflag.String("mac", "", "Virtual machine MAC address. Required if the virtual machine has multiple network cards.")
	pflag.Bool("connected", true, "Network card status. set to false to disconnect. True to connect. Default is true")

	pflag.ErrHelp = errors.New("Example: " + os.Args[0] + " --hostname='192.168.10.55' --port=9440 --username='admin' --password='supersecretpassword' --insecure=false --vmname='Windows VM test' --mac='50:6b:8d:57:a3:81' --connected=false")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	if viper.GetString("hostname") == "" || viper.GetString("username") == "" || viper.GetString("password") == "" || viper.GetString("vmname") == "" {
		pflag.Usage()
		os.Exit(1)
	}
}

func main() {
	declareFlags()

	vMNameToBeDisconnected := viper.GetString("vmname")
	macAddressToDisconnect := viper.GetString("mac")
	nicConnected := viper.GetBool("connected")
	/***********************************************************************************************************************************************/
	prismLogin := prismLoginDetails{hostname: viper.GetString("hostname"), port: viper.GetInt("port"), username: viper.GetString("username"), password: viper.GetString("password"), insecure: viper.GetBool("insecure")}
	allVMs, err := getAllVMs(&prismLogin)
	if err != nil {
		log.Fatal(err)
	}
	/***********************************************************************************************************************************************/
	activeVMUUID := getVMUUID(&allVMs, vMNameToBeDisconnected)
	activeVMNICs, err := getVMNICs(&prismLogin, activeVMUUID)
	if err != nil {
		log.Fatal(err)
	}
	if macAddressToDisconnect == "" {
		macAddressToDisconnect, err = getvmNICMAC(&activeVMNICs)
		if err != nil {
			log.Fatal(err)
		}
	}
	activeVMNetworkUUID := getVMNetworkUUID(&activeVMNICs, macAddressToDisconnect)
	/***********************************************************************************************************************************************/
	err = setNicConnectivity(&prismLogin, activeVMUUID, macAddressToDisconnect, activeVMNetworkUUID, nicConnected)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if nicConnected {
		status = "connected"
	} else {
		status = "disconnected"
	}
	fmt.Println("Network Card with MAC address \"" + macAddressToDisconnect + "\" connected to virtual machine \"" + vMNameToBeDisconnected + "\" is now " + status + ".")
	/***********************************************************************************************************************************************/
}
