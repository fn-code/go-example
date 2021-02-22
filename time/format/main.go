package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp/syntax"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"
)

var zoneDirs = map[string]string{
	"android":   "/system/usr/share/zoneinfo/",
	"darwin":    "/usr/share/zoneinfo/",
	"dragonfly": "/usr/share/zoneinfo/",
	"freebsd":   "/usr/share/zoneinfo/",
	"linux":     "/usr/share/zoneinfo/",
	"netbsd":    "/usr/share/zoneinfo/",
	"openbsd":   "/usr/share/zoneinfo/",
	// "plan9":"/adm/timezone/", -- no way to test this platform
	"solaris": "/usr/share/lib/zoneinfo/",
	"windows": `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones\`,
}

var zoneDir string

var timeZones []string

var countryTz = map[string]string{
	"Jakarta":  "Asia/Jakarta",
	"Jayapura": "Asia/Jayapura",
	"Makassar": "Asia/Makassar",
}

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC().Format(time.RFC3339))
	fmt.Println(time.Now().UTC())

	// 2021-02-15 23:50:43.45621 +0800 WITA m=+0.000103208
	times, err := time.Parse(time.RFC3339, "2021-02-15T15:50:43Z")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(times)
	fmt.Println(times.String())
	fmt.Println(time.Now().Location())
	fmt.Println(time.Now().Local().UTC().Zone())

	fmt.Println("-> converted time")

	fmt.Printf("UTC %v \n", times)
	fmt.Printf("Local %v \n", times.In(time.Now().Location()))
	fmt.Printf("Jakarta %v \n", times.In(timeIn("Jakarta")))
	fmt.Printf("Makassar %v \n", times.In(timeIn("Makassar")))
	fmt.Printf("Jayapura %v \n", times.In(timeIn("Jayapura")))

	//ListTimeZones()
}

func timeIn(name string) *time.Location {
	loc, err := time.LoadLocation(countryTz[name])
	if err != nil {
		panic(err)
	}
	return loc
}

// InSlice ... check if an element is inside a slice
func InSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// ReadTZFile ... read timezone file and append into timeZones slice
func ReadTZFile(path string) {
	files, _ := ioutil.ReadDir(zoneDir + path)
	for _, f := range files {
		if f.Name() != strings.ToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}
		if f.IsDir() {
			ReadTZFile(path + "/" + f.Name())
		} else {
			tz := (path + "/" + f.Name())[1:]
			// check if tz is already in timeZones slice
			// append if not
			if !InSlice(tz, timeZones) { // need a more efficient method...

				// convert string to rune
				tzRune, _ := utf8.DecodeRuneInString(tz[:1])

				if syntax.IsWordChar(tzRune) { // filter out entry that does not start with A-Za-z such as +VERSION
					timeZones = append(timeZones, tz)
				}
			}
		}
	}
}

func ListTimeZones() {
	if runtime.GOOS == "nacl" || runtime.GOOS == "" {
		fmt.Println("Unsupported platform")
		os.Exit(0)
	}

	// detect OS
	fmt.Println("Time zones available for : ", runtime.GOOS)
	fmt.Println("------------------------")

	fmt.Println("Retrieving time zones from : ", zoneDirs[runtime.GOOS])

	if runtime.GOOS != "windows" {
		for _, zoneDir = range zoneDirs {
			ReadTZFile("")
		}
	} else {
		// let's handle Windows
		// if you're building this on darwin/linux
		// chances are you will encounter
		// undefined: registry in registry.OpenKey error message
		// uncomment below if compiling on Windows platform

		//k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones`, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)

		//if err != nil {
		// fmt.Println(err)
		//}
		//defer k.Close()

		//names, err := k.ReadSubKeyNames(-1)
		//if err != nil {
		// fmt.Println(err)
		//}

		//fmt.Println("Number of timezones : ", len(names))
		//for i := 0; i <= len(names)-1; i++ {
		// check if tz is already in timeZones slice
		// append if not
		// if !InSlice(names[i], timeZones) { // need a more efficient method...
		//  timeZones = append(timeZones, names[i])
		// }
		//}

		// UPDATE : Reading from registry is not reliable
		// better to parse output result by "tzutil /g" command
		// REMEMBER : There is no time difference between Coordinated Universal Time and Greenwich Mean Time ....
		cmd := exec.Command("tzutil", "/l")

		data, err := cmd.Output()

		if err != nil {
			panic(err)
		}

		fmt.Println("UTC is the same as GMT")
		fmt.Println("There is no time difference between Coordinated Universal Time and Greenwich Mean Time ....")
		GMTed := bytes.Replace(data, []byte("UTC"), []byte("GMT"), -1)

		fmt.Println(string(GMTed))

	}

	now := time.Now()

	for _, v := range timeZones {

		if runtime.GOOS != "windows" {

			location, err := time.LoadLocation(v)
			if err != nil {
				fmt.Println(err)
			}

			// extract the GMT
			t := now.In(location)
			t1 := fmt.Sprintf("%s", t.Format(time.RFC822Z))
			tArray := strings.Fields(t1)
			gmtTime := strings.Join(tArray[4:], "")
			hours := gmtTime[0:3]
			minutes := gmtTime[3:]

			gmt := "GMT" + fmt.Sprintf("%s:%s", hours, minutes)
			fmt.Println(gmt + " " + v)

		} else {
			fmt.Println(v)
		}

	}
	fmt.Println("Total timezone ids : ", len(timeZones))
}
