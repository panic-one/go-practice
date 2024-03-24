package gazou

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type Gazou struct{}

func NewGazou() *Gazou {
	return &Gazou{}
}

var inputFormat = flag.String("i", "jpg", "変換前の画像の形式の指定")
var outputFormat = flag.String("o", "png", "変換後の画像の形式の指定")

func (c *Gazou) ConvertJPGToPNG(dir string, ipf string, opf string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("ディレクトリ内を走査するときにエラーが発生しました: %v", err)
		}
		if !info.IsDir() && filepath.Ext(path) == "."+*inputFormat {
			return henkan(path, *outputFormat)
		}
		return nil
	})
}

func henkan(filePath, outputFormat string) error {
	flag.Parse()
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

	if err := png.Encode(outFile, img); err != nil {
		fmt.Errorf("Encodeに失敗しました: %v", err)
	}

	return nil
}
