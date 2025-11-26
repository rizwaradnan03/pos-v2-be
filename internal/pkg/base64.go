package pkg

// import (
// 	"encoding/base64"
// 	"image/jpeg"
// 	"strings"

// 	"gocv.io/x/gocv"
// )

// func DecodeBase64ToMat(b64 string) (gocv.Mat, error) {
// 	if strings.Contains(b64, ","){
// 		parts := strings.Split(b64, ",")
// 		b64 = parts[1]
// 	}

// 	data, err := base64.StdEncoding.DecodeString(b64)
// 	if err != nil {
// 		return gocv.NewMat(), err
// 	}

// 	img, err := jpeg.Decode(strings.NewReader(string(data)))
// 	if err != nil {
// 		return gocv.NewMat(), err
// 	}

// 	buf, err := gocv.ImageToMatRGB(img)
// 	if err != nil {
// 		return gocv.NewMat(), err
// 	}

// 	return buf, nil
// }
