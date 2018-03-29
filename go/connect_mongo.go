package main

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type SpecialType struct {
	ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
	SpecialType    string        `bson:"specialType" json:"specialType"`
	Description    string        `bson:"description" json:"description"`
	ContractPeriod int           `bson:"contractPeriod" json:"contractPeriod"`
	ActiveFlag     string        `bson:"activeFlag" json:"activeFlag"`
	AdvancePayment int           `bson:"advancePayment" json:"advancePayment, omitempty"`
}

func connectMongoDB() *mgo.Session {
	session, err := mgo.Dial("mongodb://omrdev:omrdev2014@andromeda11.tac.co.th:27017")
	if err != nil {
		fmt.Println("connect error")
		fmt.Println(err)
	}

	return session
}

func findOne() {
	session := connectMongoDB()
	c := session.DB("dev2").C("SpecialTypeMappingContractPeriod")
	var specialTypes SpecialType
	err := c.Find(bson.M{"specialType": "C11"}).One(&specialTypes)
	if err != nil {
		fmt.Println("find one error")
		fmt.Println(err)
	}

	fmt.Printf("\n%#v\n", specialTypes)
}

func findAll() {
	session := connectMongoDB()
	c := session.DB("dev2").C("SpecialTypeMappingContractPeriod")
	var specialTypes []SpecialType
	err := c.Find(bson.M{"specialType": "C11"}).All(&specialTypes)
	if err != nil {
		fmt.Println("find apll error")
		fmt.Println(err)
	}

	fmt.Printf("\n%d\n", len(specialTypes))
}

func insert() {
	session := connectMongoDB()
	c := session.DB("dev2").C("SpecialTypeMappingContractPeriod")
	var specialTypes SpecialType
	specialTypes.ID = bson.NewObjectId()
	specialTypes.SpecialType = "TEST"
	specialTypes.Description = "Test for insert"
	specialTypes.ActiveFlag = "A"
	specialTypes.ContractPeriod = 10
	specialTypes.AdvancePayment = 6000

	err := c.Insert(specialTypes)
	if err != nil {
		fmt.Println("insert error")
		fmt.Println(err)
	}
}

func delete() {
	session := connectMongoDB()
	c := session.DB("dev2").C("SpecialTypeMappingContractPeriod")
	err := c.Remove(bson.M{"specialType": "TEST"})
	if err != nil {
		fmt.Println("delete error")
		fmt.Println(err)
	}
}

func update() {
	session := connectMongoDB()
	c := session.DB("dev2").C("SpecialTypeMappingContractPeriod")
	var specialTypes SpecialType
	specialTypes.ID = bson.ObjectIdHex("5abc895e2bfc1748cb9da2a3")
	specialTypes.SpecialType = "TEST2"
	specialTypes.Description = "Test for insert2"
	specialTypes.ActiveFlag = "A"
	specialTypes.ContractPeriod = 10
	specialTypes.AdvancePayment = 6000
	err := c.Update(bson.M{"specialType": "TEST"}, specialTypes)
	if err != nil {
		fmt.Println("update error")
		fmt.Println(err)
	}
}

func main() {
	findOne()
	findAll()
	insert()
	delete()
	update()
}
