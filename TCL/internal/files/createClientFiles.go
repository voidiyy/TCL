package files

import (
	"Hook_TCL/internal/logger"
	"log"
	"os"
)

func CreateClientDirectories(clientName string) (string, string, error) {

	var (
		err  error
		path = "clientSessions/" + clientName
	)

	err = os.MkdirAll(path, 0755)
	if err != nil {
		logger.Global().Error("error creating client directory")
		return "", "", err
	}

	tdlibDBDir := path + "/tdlib_db_" + clientName
	err = os.MkdirAll(tdlibDBDir, 0755)
	if err != nil {
		log.Printf("Failed to create client DB directory: %v", err)
		return "", "", err
	}

	tdlibFilesDir := path + "/tdlib_files_" + clientName
	err = os.MkdirAll(tdlibFilesDir, 0755)
	if err != nil {
		log.Printf("Failed to create client files directory: %v", err)
		return "", "", err
	}

	return tdlibDBDir, tdlibFilesDir, nil
}
