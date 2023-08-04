package main

import (
	p "SpringMassClustering/proto"
	"errors"
	"fmt"
	pb "google.golang.org/protobuf/proto"
	ini "gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	cfg, err := ini.Load("settings.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	DT, err := cfg.Section("").Key("DT").Float64()
	if err != nil {
		log.Println(fmt.Errorf("failed to load DT, defaulting to 0.1"))
		DT = 0.1
	}

	gens, err := cfg.Section("").Key("Generations").Int()
	if err != nil {
		log.Println(fmt.Errorf("failed to load generations, defaulting to 10000"))
		gens = 10000
	}

	in, err := ioutil.ReadFile("bin/masses.bin")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	massesPB := &p.Masses{}
	if err := pb.Unmarshal(in, massesPB); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	in, err = ioutil.ReadFile("bin/springs.bin")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	springsPB := &p.Springs{}
	if err := pb.Unmarshal(in, springsPB); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	springs := []*Spring{}
	masses := []*Mass{}

	for i := 0; i < len(massesPB.GetMasses()); i++ {
		mass := NewMassFromPB(massesPB.GetMasses()[i])
		// TODO: could be read from file potentially
		if i == 0 {
			mass.Fixed = true
		}
		masses = append(masses, mass)
	}

	searchMass := func(index int32) (*Mass, error) {
		for j := 0; j < len(masses); j++ {
			if masses[j].Id == index {
				return masses[j], nil
			}
		}

		return nil, errors.New("couldn't find mass")
	}

	for i := 0; i < len(springsPB.GetSprings()); i++ {
		spring := NewSpringFromPB(springsPB.GetSprings()[i])

		spring.Mass1, err = searchMass(springsPB.GetSprings()[i].M1)
		if err != nil {
			log.Println(err)
			return
		}

		spring.Mass2, err = searchMass(springsPB.GetSprings()[i].M2)
		if err != nil {
			log.Println(err)
			return
		}

		springs = append(springs, spring)
	}

	trajectories := p.Trajectories{Trajectories: []*p.Trajectory{}}

	for i := 0; i < len(masses); i++ {
		trajectories.Trajectories = append(trajectories.Trajectories, &p.Trajectory{
			MassId: masses[i].Id,
			Points: []*p.Point{},
		})
	}

	counter := 0

	// Start simulation
	for i := 0; i < gens; i++ {
		for j := 0; j < len(springs); j++ {
			springs[j].Update()
		}

		for j := 0; j < len(masses); j++ {
			masses[j].Update(float32(DT))
			masses[j].Force = []float32{0, 0}

			trajectories.Trajectories[j].Points = append(trajectories.Trajectories[j].Points, &p.Point{
				PosX: masses[j].Position[0],
				PosY: masses[j].Position[1],
			})
		}

		if i%1000 == 0 {
			fmt.Println("Generation: ", i)

			counter += 1

			out, err := pb.Marshal(&trajectories)
			if err != nil {
				log.Fatalln("Failed to encode address book:", err)
			}

			if err := ioutil.WriteFile(fmt.Sprintf("bin/trajectories_%06d.bin", counter), out, 0644); err != nil {
				log.Fatalln("Failed to write address book:", err)
			}

			trajectories = p.Trajectories{Trajectories: []*p.Trajectory{}}

			for i := 0; i < len(masses); i++ {
				trajectories.Trajectories = append(trajectories.Trajectories, &p.Trajectory{
					MassId: masses[i].Id,
					Points: []*p.Point{},
				})
			}
		}
	}

	counter += 1

	out, err := pb.Marshal(&trajectories)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fmt.Sprintf("bin/trajectories_%06d.bin", counter), out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
