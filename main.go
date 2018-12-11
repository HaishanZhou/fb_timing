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

// https://www.tldp.org/HOWTO/html_single/Framebuffer-HOWTO/#AEN1274
// Modeline "1280x1024" DCF HR SH1 SH2 HFL VR SV1 SV2 VFL
// newmode: clock hdisp hsyncstart hsyncend  htotal  vdisp vsyncstart vsyncend vtotal
//    left_margin = HFL - SH2
//    right_margin = SH1 - HR
//    hsync_len = SH2 - SH1
//    upper_margin = VFL - SV2
//    lower_margin = SV1 - VR
//    vsync_len = SV2 - SV1
//
//    HR = xres
//    SH1 = HR + right_margin
//    SH2 = SH1 + hsync_len
//    HFL = SH2 + left_margin
//    VR = yres
//    SV1 = VR + lower_margin
//    SV2 = SV1 + vsync_len
//    VFL = SV2 + uppper_margin
//
func xorgModeline(t *timing) {
	hr := t.XRes
	sh1 := hr + t.RightMargin
	sh2 := sh1 + t.HSyncLen
	hfl := sh2 + t.LeftMargin

	vr := t.YRes
	sv1 := vr + t.LowerMargin
	sv2 := sv1 + t.VSyncLen
	vfl := sv2 + t.UpperMargin

	fmt.Println("Xorg Modeline:", t.DotClock, hr, sh1, sh2, hfl, vr, sv1, sv2, vfl)
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
	xorgModeline(t)
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
