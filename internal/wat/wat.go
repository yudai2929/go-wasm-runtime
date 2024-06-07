package wat

import (
	"fmt"
	"os/exec"

	"github.com/yudai2929/go-wasm-runtime/pkg/file"
)

func createWasmFrom(watFile string) (wasmFile string, err error) {
	wasmFile = watFile[:len(watFile)-len("wat")] + "wasm"

	// wat2wasmコマンドを実行
	cmd := exec.Command("wat2wasm", watFile, "-o", wasmFile)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to convert WAT to WASM: %v", err)
	}

	return
}

func ParseFile(watFile string) (wasmBinary []byte, err error) {
	// WASMファイルの生成
	wasmFile, wasmErr := createWasmFrom(watFile)
	if wasmErr != nil {
		return nil, wasmErr
	}

	defer func() {
		// delete the file
		if wasmErr == nil {
			_ = file.Delete(wasmFile)

		}
	}()

	// read the file
	wasmBinary, err = file.Read(wasmFile)
	if err != nil {
		return nil, err
	}

	return wasmBinary, nil
}

func ParseStr(wat string) (wasmBinary []byte, err error) {
	// WASMファイルの生成
	watFile := "temp.wat"
	createErr := file.Create(watFile, []byte(wat))
	if createErr != nil {
		return nil, createErr
	}

	defer func() {
		// delete the file
		if createErr == nil {
			_ = file.Delete(watFile)
		}
	}()

	// read the file
	wasmBinary, err = ParseFile(watFile)
	if err != nil {
		return nil, err
	}

	return wasmBinary, nil
}
