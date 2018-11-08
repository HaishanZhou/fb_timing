package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type timing struct {
	DotClock    int `json:"DotClock"`
	XRes        int `json:"XRes"`
	YRes        int `json:"YRes"`
	RightMargin int `json:"RightMargin"`
	LeftMargin  int `json:"LeftMargin"`
	HSyncLen    int `json:"HSyncLen"`
	LowerMargin int `json:"LowerMargin"`
	UpperMargin int `json:"UpperMargin"`
	VSyncLen    int `json:"VSyncLen"`
}

func calculate(t *timing) {
	scanLineTime := (t.XRes + t.LeftMargin + t.RightMargin + t.HSyncLen)
	frameTime := scanLineTime * (t.YRes + t.UpperMargin + t.LowerMargin + t.VSyncLen)
	fps := t.DotClock / frameTime

	fmt.Printf("Resolution: %dx%d\n", t.XRes, t.YRes)
	fmt.Printf("Timing: freq: %d\n", t.DotClock)
	fmt.Printf("        Right Maring (hfrontporch): \t%d\n", t.RightMargin)
	fmt.Printf("        Left Maring (hbackporch): \t%d\n", t.LeftMargin)
	fmt.Printf("        H Sync Len (hsyncwidth): \t%d\n", t.HSyncLen)
	fmt.Printf("        Lower Maring (vfrontporch): \t%d\n", t.LowerMargin)
	fmt.Printf("        Upper Maring (vbackporch): \t%d\n", t.UpperMargin)
	fmt.Printf("        V Sync Len (vsyncwidth): \t%d\n", t.VSyncLen)
	fmt.Println("fps:", fps, "Hz")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: need timing json file")
		return
	}
	timingFile := os.Args[1]

	jsonFile, err := os.Open(timingFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	var t timing
	err = json.Unmarshal(byteValue, &t)
	if err != nil {
		fmt.Println(err)
		return
	}

	calculate(&t)
}
