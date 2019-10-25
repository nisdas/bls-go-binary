package bls

import (
	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"io/ioutil"
	"testing"
)

func testUncompressedG1(t *testing.T, gen1 *G1) {
	filePath, err := bazel.Runfile("tests/g1_uncompressed_valid_test_vectors.dat")
	if err != nil {
		t.Fatal(err)
	}
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile")
	}
	var p1, p2 G1
	for i := 0; i < 1000; i++ {
		if p1.DeserializeUncompressed(buf[i*96:(i+1)*96]) != nil {
			t.Fatalf("i=%d X.Deserialize", i)
		}
		if !p1.IsEqual(&p2) {
			t.Fatalf("i=%d p1=%x\np2=%x\n", i, p1.Serialize(), p2.Serialize())
		}
		G1Add(&p2, &p2, gen1)
	}
}

func testCompressedG1(t *testing.T, gen1 *G1) {
	filePath, err := bazel.Runfile("tests/g1_compressed_valid_test_vectors.dat")
	if err != nil {
		t.Fatal(err)
	}
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile")
	}
	var p1, p2 G1
	for i := 0; i < 1000; i++ {
		if p1.Deserialize(buf[i*48:(i+1)*48]) != nil {
			t.Fatalf("err i=%d\n", i)
		}
		if !p1.IsEqual(&p2) {
			t.Fatalf("p1=%x\np2=%x\n", p1.Serialize(), p2.Serialize())
		}
		G1Add(&p2, &p2, gen1)
	}
}

func testUncompressedG2(t *testing.T, gen2 *G2) {
	filePath, err := bazel.Runfile("tests/g2_uncompressed_valid_test_vectors.dat")
	if err != nil {
		t.Fatal(err)
	}
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile")
	}
	var p1, p2 G2
	for i := 0; i < 1000; i++ {
		if p1.DeserializeUncompressed(buf[i*192:(i+1)*192]) != nil {
			t.Fatalf("i=%d X.Deserialize", i)
		}
		if !p1.IsEqual(&p2) {
			t.Fatalf("i=%d p1=%x\np2=%x\n", i, p1.Serialize(), p2.Serialize())
		}
		G2Add(&p2, &p2, gen2)
	}
}

func testCompressedG2(t *testing.T, gen2 *G2) {
	filePath, err := bazel.Runfile("tests/g2_compressed_valid_test_vectors.dat")
	if err != nil {
		t.Fatal(err)
	}
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile")
	}
	var p1, p2 G2
	for i := 0; i < 1000; i++ {
		if p1.Deserialize(buf[i*96:(i+1)*96]) != nil {
			t.Fatalf("err i=%d\n", i)
		}
		if !p1.IsEqual(&p2) {
			t.Fatalf("p1=%x\np2=%x\n", p1.Serialize(), p2.Serialize())
		}
		G2Add(&p2, &p2, gen2)
	}
}

func Test(t *testing.T) {
	if Init(BLS12_381) != nil {
		t.Fatalf("Init")
	}
	var gen1 G1
	if gen1.SetString("1 0x17f1d3a73197d7942695638c4fa9ac0fc3688c4f9774b905a14e3a3f171bac586c55e83ff97a1aeffb3af00adb22c6bb 0x08b3f481e3aaa0f1a09e30ed741d8ae4fcf5e095d5d00af600db18cb2c04b3edd03cc744a2888ae40caa232946c5e7e1", 16) != nil {
		t.Fatalf("gen1.SetString")
	}
	var gen2 G2
	if gen2.SetString("1 0x024aa2b2f08f0a91260805272dc51051c6e47ad4fa403b02b4510b647ae3d1770bac0326a805bbefd48056c8c121bdb8 0x13e02b6052719f607dacd3a088274f65596bd0d09920b61ab5da61bbdc7f5049334cf11213945d57e5ac7d055d042b7e 0x0ce5d527727d6e118cc9cdc6da2e351aadfd9baa8cbdd3a76d429a695160d12c923ac9cc3baca289e193548608b82801 0x0606c4a02ea734cc32acd2b02bc28b99cb3e287e85a763af267492ab572e99ab3f370d275cec1da1aaa9075ff05f79be", 16) != nil {
		t.Fatalf("gen2.SetString")
	}

	testUncompressedG1(t, &gen1)
	testCompressedG1(t, &gen1)
	testUncompressedG2(t, &gen2)
	testCompressedG2(t, &gen2)
}
