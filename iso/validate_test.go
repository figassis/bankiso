package iso

import (
	"io/ioutil"
	"os"
	"testing"
)

var gopath = os.Getenv("GOPATH")

//@TODO: perform tests for invalid XML as well
func TestMakeISO20022(t *testing.T) {
	m := make(map[string]string)
	m["pain.009.001.04"] = gopath + "/src/github.com/figassis/bankiso/msg-examples/payments/pain.009.001.04/Business example1 pain.009.001.04.xml"
	m["pain.010.001.04"] = gopath + "/src/github.com/figassis/bankiso/msg-examples/payments/pain.010.001.04/Business example1 pain.010.001.04.xml"
	m["pain.011.001.04"] = gopath + "/src/github.com/figassis/bankiso/msg-examples/payments/pain.011.001.04/Business example1 pain.011.001.04.xml"
	m["pain.012.001.04"] = gopath + "/src/github.com/figassis/bankiso/msg-examples/payments/pain.012.001.04/Business example1 pain.012.001.04.xml"

	for key, value := range m {
		data, err := ioutil.ReadFile(value)
		if err != nil {
			t.Error("TestMakeISO20022 - Cannot read file: ", err)
		}

		if _, err = makeISO20022(key, string(data)); err != nil {
			t.Error("TestMakeISO20022 - Failed to build from valid iso20022 message: ", err)
		}
	}

}

func TestMakeISO8583(t *testing.T) {
}

func TestIsXML(t *testing.T) {
}

func TestValidateISO20022(t *testing.T) {
}

func TestValidateISO8583(t *testing.T) {
}

func TestValidateISO(t *testing.T) {
}

func TestDecode(t *testing.T) {
}
