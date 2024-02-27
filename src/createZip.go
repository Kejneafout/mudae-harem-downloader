package src

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func CreateZip(exportsDirectory, jsonFilename, imagesDirectory string) error {
	// Check if imagesDirectory exists, if not, create it
	if _, err := os.Stat(exportsDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(exportsDirectory, os.ModePerm); err != nil {
			return err
		}
	}

	// Open JSON file
	jsonFile, err := os.Open(jsonFilename)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// Create zip file
	zipFilename := filepath.Join(exportsDirectory, fmt.Sprintf("export_%s.zip", time.Now().Format("20060102_150405")))
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add JSON file to zip
	jsonFileInfo, err := jsonFile.Stat()
	if err != nil {
		return err
	}
	jsonHeader, err := zip.FileInfoHeader(jsonFileInfo)
	if err != nil {
		return err
	}
	jsonHeader.Name = jsonFilename
	jsonFileInZip, err := zipWriter.CreateHeader(jsonHeader)
	if err != nil {
		return err
	}
	if _, err := io.Copy(jsonFileInZip, jsonFile); err != nil {
		return err
	}

	// Add images directory to zip
	err = filepath.Walk(imagesDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(imagesDirectory, path)
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = filepath.Join(imagesDirectory, relPath)
		if info.IsDir() {
			header.Name += "/"
		}
		fileInZip, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			if _, err := io.Copy(fileInZip, file); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Printf("Zip file %s created successfully.\n", zipFilename)

	// Delete data.json and images/ after .zip creation
	if err := os.RemoveAll(imagesDirectory); err != nil {
		return err
	}
	if err := os.Remove(jsonFilename); err != nil {
		return err
	}

	return nil
}
