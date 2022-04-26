package main

import (
	
	
	"Skyplatform/cmd"
	"github.com/pkg/profile"
	
)

func main() {

    
	if version.Profile == "on"{
		if version.ProfileType == "CPU"{
			defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
		}
		if version.ProfileType == "MEM"{
			defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()
		}
	}

    cmd.Execute()
	
}
