package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tomba-hopkins/tools-kecil-kecilan/nofap-app-cli-golang/pkg/app"
)

func main() {

	lokasiFile := "../data/days.json"

	data, err := app.ReadFileData(lokasiFile)
	if err != nil {
		log.Fatalf("Ada error pokoknya mah: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Beneran berhasil kan? (y/n)")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "y" {
		data.Days++
		fmt.Println(app.QuotesBiarSemangat())
		fmt.Println("Streak mu adalah : ", data.Days)
		err = app.WriteFileData(lokasiFile, data)
		if err != nil {
			log.Fatalf("Error lah pokoknya: %v", err)
		}
	} else if input == "n"{
		fmt.Println("Kocag lu out aja lah out")
		os.Exit(0)
	} else {
		fmt.Println("Kocag jawab ngawur dahlah")
		os.Exit(1)
	}

	fmt.Println("Enter biar out")
	fmt.Scanln()


}