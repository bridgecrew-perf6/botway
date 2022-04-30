package fleet

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/tools/templates/telegram/rust"
	"github.com/abdfnx/looker"
)

func TelegramRustFleet(botName string) {
	_, err := looker.LookPath("rust")
	fleetPath, ferr := looker.LookPath("fleet")

	if err != nil {
		log.Fatal("error: rust is not installed")
	} else if ferr != nil {
		log.Fatal("error: fleet is not installed")
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.rs"), []byte(rust.MainRsContent()), 0644)
		cargoFile := os.WriteFile(filepath.Join(botName, "Cargo.toml"), []byte(rust.CargoFileContent(botName)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName)), 0644)
		procFile := os.WriteFile(filepath.Join(botName, "Procfile"), []byte("process: ./" + botName), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(rust.Resources()), 0644)

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if cargoFile != nil {
			log.Fatal(cargoFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if procFile != nil {
			log.Fatal(procFile)
		}

		rustUpPath, err := looker.LookPath("rustup")

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		rustUpCmd := rustUpPath + " default nightly"

		rustUp := exec.Command("bash", "-c", rustUpCmd)

		if runtime.GOOS == "windows" {
			rustUp = exec.Command("powershell.exe", rustUpCmd)
		}

		rustUp.Dir = botName
		rustUp.Stdin = os.Stdin
		rustUp.Stdout = os.Stdout
		rustUp.Stderr = os.Stderr
		err = rustUp.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		fleetBuild := fleetPath + " build"

		buildCmd := exec.Command("bash", "-c", fleetBuild)

		if runtime.GOOS == "windows" {
			buildCmd = exec.Command("powershell.exe", fleetBuild)
		}

		buildCmd.Dir = botName
		buildCmd.Stdin = os.Stdin
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		err = buildCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
}