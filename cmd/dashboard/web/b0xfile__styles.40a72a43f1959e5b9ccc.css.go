// Code generaTed by fileb0x at "2021-12-07 10:11:10.6017344 +0800 CST m=+0.177909301" from config file "b0x.yaml" DO NOT EDIT.
// modified(2021-12-07 10:11:05.7530359 +0800 CST)
// original path: frontend\dist\styles.40a72a43f1959e5b9ccc.css

package web

import (
	"os"
)

// FileStyles40a72a43f1959e5b9cccCSS is "/styles.40a72a43f1959e5b9ccc.css"

func init() {

	f, err := FS.OpenFile(CTX, "/styles.40a72a43f1959e5b9ccc.css", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileStyles40a72a43f1959e5b9cccCSS)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}