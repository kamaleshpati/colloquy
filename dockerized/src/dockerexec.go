package dockerized

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func DockerRunner() {
	dockerUp()
	handleSystemSignal(dockerDown)
}

func handleSystemSignal(dockerDown func()) (res int) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSEGV, os.Interrupt)
	for {
		s := <-signalChannel
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, os.Interrupt:
			dockerDown()
			return 0
		case syscall.SIGHUP:
		case syscall.SIGSEGV:
		default:
			return 0
		}
	}
}

func dockerUp() {
	cmd := exec.Command("docker-compose", "-f", "dockerized/resources/compose.test.yml", "up")
	errorLogger(cmd)
}

func dockerDown() {
	cmd := exec.Command("docker-compose", "-f", "./dockerized/resources/compose.test.yml", "down", "--volumes")
	errorLogger(cmd)
}

func errorLogger(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
