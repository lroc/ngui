package ngui

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"path"
)

const (
	manifest_filename  = `manifest.json`
	first_page         = `first_page`
	application_title  = `application_title`
	launch_width       = `launch_width`
	launch_height      = `launch_height`
	enable_transparent = `enable_transparent`
)

type Manifest struct {
	manifest revManifest
}

type revManifest map[string]interface{}

func (a *Manifest) FirstPage() string {
	return a.Get(first_page).(string)
}

func (a *Manifest) ApplicationTitle() string {
	return a.Get(application_title).(string)
}

func (a *Manifest) EnableTransparent() bool {
	return a.Get(enable_transparent).(bool)
}

func (a *Manifest) LaunchWidth() int32 {
	return int32(a.Get(launch_width).(float64))
}

func (a *Manifest) LaunchHeight() int32 {
	return int32(a.Get(launch_height).(float64))
}

func (a *Manifest) Get(key string) interface{} {
	v, _ := a.manifest[key]
	return v
}

func (a *Manifest) Load() *Manifest {
	manifestPath := path.Join(a.Path(), manifest_filename)
	//fmt.Printf("manifestPath=%v\n", manifestPath)
	data, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		panic("Load Manifest")
	}
	//fmt.Println(err)
	json.Unmarshal(data, &a.manifest)
	return a
}

func (a Manifest) Path() string {
	return ExePath()
}