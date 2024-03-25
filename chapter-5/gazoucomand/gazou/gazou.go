package gazou

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type Gazou struct{}

func NewGazou() *Gazou {
	return &Gazou{}
}

func (c *Gazou) Convert(dir string, ipf string, opf string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("ディレクトリ内を走査するときにエラーが発生しました: %v", err)
		}
		if !info.IsDir() && filepath.Ext(path) == ipf {
			return henkan(path, ipf, opf)
		}
		return nil
	})
}

func henkan(filePath, inputFormat string, outputFormat string) error {
	sf, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ファイル読み込みに失敗しました", err)
	}
	defer sf.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		fmt.Errorf("Decodeに失敗しました: %v", err)
	}

	outputPath := filePath[:len(filePath)-len(filepath.Ext(filePath))] + "." + outputFormat
	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Errorf("PNGファイルの作成に失敗しました: %v", err)
	}

	defer outFile.Close()

	if inputFormat == "jpeg" {
		if err := png.Encode(outFile, img); err != nil {
			fmt.Errorf("Encodeに失敗しました: %v", err)
		}
	} else if inputFormat == "png" {
		if err := jpeg.Encode(outFile, img, nil); err != nil {
			fmt.Errorf("Encodeに失敗しました: %v", err)
		}
	} else {
		fmt.Errorf("invalid file type")
	}

	return err
}
