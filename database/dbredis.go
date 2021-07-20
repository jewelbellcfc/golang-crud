package database

import (
	"github.com/go-redis/redis"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"log"
	"strconv"
	"strings"
)
type Person struct {
	ID string `json:"id"`
	FullName string `json:"full_name"`
	Gender string `json:"gender"`
}
func conn() *redis.Client{
	dbr := redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		Password:           "",
		DB:                 0,
	})
	err := dbr.Ping().Err()
	if err != nil {
		log.Fatal( "redis connect fail", err )
	}

	return dbr
}
const IDCounter = "personid"
const ListIdSet = "listid"
const gender = "updating"
//tao person
func  CreatePerson(fullname string )  {
	dbr := conn()
	defer dbr.Close()
//tao key
	id,err:= dbr.Incr(IDCounter).Result()
	if err != nil {
		log.Fatal("failed to increment id counter",err)
	}
	personid := "key:" + strconv.Itoa(int(id))
//add key vao list

	err = dbr.SAdd(ListIdSet,personid).Err()
	if err !=nil{
		log.Fatal(" set failed",err)
	}
	//add key-value
	ps := map[string]interface{}{"fullname": fullname, "gender": gender}
	err = dbr.HMSet(personid,ps).Err()
	if err != nil{
			log.Fatal("save failed",err)
	}
	fmt.Println("succeess")
}
//get all list
func ListPerson() []Person {
	dbr := conn()
	defer dbr.Close()
	HashName,err := dbr.SMembers(ListIdSet).Result()
	if err != nil{
		log.Fatal("get all ID failed",err)
	}
	listps := []Person{}
	for _ , HashName := range HashName {
		id:= strings.Split(HashName,":")[1]
		personMap,err := dbr.HGetAll(HashName).Result()
		if err !=nil{
			log.Fatal("get person failed from %s - %v\n",HashName,err)
		}
		var ps Person
		ps= Person{id, personMap["fullname"], personMap["gender"]}
		listps =append(listps,ps)
	}
	if len(listps)==0{
		fmt.Println("no person found")
		return nil
	}
	return listps
}
//delete person
func DeletePerson(id string){
	dbr:= conn()
	defer dbr.Close()
	n,err := dbr.Del("key:"+id).Result()
	if err != nil{
		log.Fatal("delete key fail %s - %v \n",id,err)
	}
	if n>0 {
		err = dbr.SRem(ListIdSet,"key:"+id).Err()
		if err!=nil {
			log.Fatal("delete key fail %s - %v \n",id,err)
		}
		fmt.Println("delete success")
	} else {
		fmt.Println("id not found",id)
	}
}
// update
func UpdatePerson(id,fullname,gender string){
	dbr := conn()
	defer dbr.Close()
	exists,err := dbr.SIsMember(ListIdSet,"key:"+id).Result()
	if err != nil{
		log.Fatal("cannot find id %s exist %v",id,err)
	}
	if !exists {
		log.Fatal("id not exist %s \n",id)
	}
	updatePerson := map[string]interface{}{}
	updatePerson["fullname"] = fullname
	updatePerson["gender"] = gender
	err = dbr.HMSet("todo:"+id,updatePerson).Err()
	if err!= nil{
		log.Fatal("update failed")
	}
	fmt.Println("success")
}