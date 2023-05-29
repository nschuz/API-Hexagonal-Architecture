package postgres

import (
	"fmt"
	"sync"
	"time"

	"github.com/nschuz/go-arquitectura-hexagonal/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

//PATRON SINGLETON UNA MULTIPLE CONEXION

var onceDBLoad sync.Once

var tables = []interface{}{
	&entity.User{},
}

func connect() *gorm.DB {
	//si exsite una conexion se ocupe y no crea multiples conexiones
	//la funcion solo se ejcuta una solavez
	onceDBLoad.Do(func() {
		source := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s",
			"127.0.0.1",
			"postgres",
			"postgres",
			"app",
			"5433",
		)

		var i int //sabe el numero de intentos
		//vamos hacer intentos de conexion
		for {
			var err error
			if i >= 30 {
				panic("Failed to connect: " + source)
			}
			time.Sleep(3 * time.Second) //intena cada tres segundos
			db, err = gorm.Open(postgres.Open(source), &gorm.Config{})

			//db es el  apuntador  de la instacia de la base de datos
			if err != nil {
				log.Info("Retitub connection...", err)
				i++
				continue
			}
			break
		}
		migrate()
		log.Info("Connection established to database")

	})
	return db
}

func migrate() {
	for _, table := range tables {
		db.AutoMigrate(table)
	}
}
