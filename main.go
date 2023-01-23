package main

import (
	"log"
	"net/http"
	"os"
	"persons/models"
	"persons/storage"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type person struct{
	Name 	  string	 `json:"name"`
	Num       int32		 `json:"num"`
	Gender    string	 `json:"gender"`
	BloodType string     `json:"bloodtype"`
}


type repository struct{
	DB *gorm.DB
}

func main(){


	router:=gin.Default()
	router.POST("/CreateUser", r.CreateUser)
	router.DELETE("/DeleteUser/:Id", r.DeleteUser)
	router.GET("/GetUser/:Id", r.GetUser)
	router.GET("/GetPerson/:Id", r.GetPerson)
	router.RUN("localhost:8080")


	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal(err)
	}

	config:= &storage.Config{
		Host:os.Getenv("DB_HOST"),
		Port:os.Getenv("DB_PORT"),
		Password:os.Getenv("DB_PASSWORD"),
		User:os.Getenv("DB_USER"),
		DBname:os.Getenv("DB_DBNAME"),
	}
	db,err1=Storage.Newconnection(config)
	if err1!=nil{
		log.Fatal("Can't read the database")
	}
	err:=models.MigratePersons(db)
	if err!=nil{
		log.Fatal("Couldn't migrate db")
	}
	r := repository{
		DB = db,
	}

}



func (r *repository)CreateUser(c *gin.Context)error{
	var NewUser person

    if err := c.BindJSON(&NewUser); err != nil {
        return err
    }
	_, err=DB.Exec("INSERT INTO books (Name, Num, Gender, BloodType) VALUES ($1, $2, $3, $4)",NewUserk.name, NewUser.Num, NewUser.Gender, NewUser.BloodType)
	if err!=nil{
		c.JSON(400,gin.H{"Message":"entry isn't posted"})
		return err
	}
	c.JSON(200,gin.H{"Message":"User deleted successfully","data":personModels})		
	return nil
}


func (r *repository)DeleteUser(c *gin.Context)error{
	personModels := &[]models.person{}
	id := c.Param("id")
	if id ==""{
		c.JSON(400,gin.H{"Message":"id can't be empty"})
	return nil
	}
	if err:=r.DB.Delete(personModels,id).Error; err!=nil{
		c.JSON(400,gin.H{"Message":"Can't delete"})
		return err
	}
	c.JSON(200,gin.H{"Message":"User deleted successfully","data":personModels})
	return nil

}


func (r *repository)GetUser(c *gin.Context)error{
	personModels := &[]models.person{}
	id := c.Param("id")
	if id ==""{
		c.JSON(400,gin.H{"Message":"id can't be empty"})
	return nil
	}

	err:=r.DB.Where("id=?",id)First(personModels).Error
	if err!=nil{
		c.JSON(400,gin.H{"Message":err})
		return err
	}
	c.JSON(200,gin.H{"Message":"User fetched successfully","data":personModels})
	return nil
}

func (r *repository) GetPerson(c *gin.Context)error{
	personModels := &[]models.person{}
	
	err:= r.DB.Find(personModels).Error
	if err!=nil{
		c.JSON(400,gin.H{"Message":err})
		return err
	}
	c.JSON(200,gin.H{"Message":"Person fetched successfully","data":personModels})
	return nil
​}

