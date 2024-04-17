package testing

import (
	"blockchains/src/blocks"
	"testing"
)


func TestIndexToHex(t *testing.T) {
	var x int = 1
	genX := blocks.IntToHex(int64(x))
	if genX == nil {
		t.Errorf("Generate Hex nil\n")
	}
	t.Logf("GenX: %v\n", genX)
}

func TestValidate(t *testing.T) {
	genBlock := blocks.NewBlock("this is first", []byte{})
	newPow := blocks.NewProofOfWork(genBlock)
	if newPow == nil {
		t.Error("new pow is nil")
	}
	if !newPow.Validate() {
		t.Error("Can not validate new pow")
	}


}