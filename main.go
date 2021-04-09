package main

import (
	"log"

	"github.com/mehre/go-design-pattern/builder"
	"github.com/mehre/go-design-pattern/facade"
	"github.com/mehre/go-design-pattern/factory"
	"github.com/mehre/go-design-pattern/objectpool"
	"github.com/mehre/go-design-pattern/repository"
	"github.com/mehre/go-design-pattern/singleton"
)

func main() {
	singletonRacing()
}

func singletonRacing() {
	for i := 0; i < 10; i++ {
		go singleton.GetDBInstance()
	}
}

func singletonSync() {
	for i := 0; i < 10; i++ {
		go singleton.GetDBInstanceSync()
	}
}

func builderBackground() {
	computer1 := builder.NewComputer("intel", "corsair", "wd", "nvidia", "LG")
	computer2 := builder.NewComputer("intel", "v-gen", "", "", "")

	log.Println("computer 1 = ", computer1)
	log.Println("computer 2 = ", computer2)
}

func builderImplementation() {
	computer1 := builder.NewComputerBuilder().
		SetProcessor("intel").
		SetRam("corsair").
		SetHarddisk("wd").
		SetVga("nvidia").
		SetMonitor("lg").
		Build()

	computer2 := builder.NewComputerBuilder().
		SetProcessor("intel").
		SetRam("v-gen").
		Build()

	log.Println("computer 1 = ", computer1)
	log.Println("computer 2 = ", computer2)
}

func factoryEngine() {
	carEngine, _ := factory.GetEngine(factory.Car)
	carEngine.Assemble()

	list := []int{factory.Car, factory.Motorcycle}

	for _, engineType := range list {
		engine, err := factory.GetEngine(engineType)
		if err != nil {
			continue
		}
		factory.AssemblingEngine(engine)
	}
}

func GetDBPooling() {
	objectpool.InitDB()
	log.Println(objectpool.GetDBConnection())
}

func repositoryPattern() {
	user := repository.User{
		ID:   "777",
		Name: "Bob",
	}

	var repository repository.UserRepository

	repository.Insert(user)
	repository.Update(user)
	repository.Delete(user.ID)
}

func FacadePattern() {
	var bankFacade facade.BankFacade

	bankFacade.Transfer("001", "002", 50000)
}
