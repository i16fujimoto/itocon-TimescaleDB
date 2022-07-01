package main

import (
   "context"
   "fmt"
   "os"
   "time"
   "net/http"

	"github.com/gin-gonic/gin"
   //"github.com/jackc/pgx/v4"
   "github.com/jackc/pgx/v4/pgxpool"
)

type Sensor struct {
	sensor_id string `json:"sensor_id"`
	sensor_name string `json:"sensor_name"`
	user_id string  `json:"user_id"`
	is_active bool `json:"is_active"`
	created_at time.Time `json:"created_at"`
}

func setRouter(ctx context.Context, dbpool *pgxpool.Pool) *gin.Engine {

	// gin.Default()は *gin.Engineを返す関数
	// *gin.Engineは「エンドポイントの追加」や「ミドルウェアの登録」をおこなってくれるもの
	r := gin.Default()

	// 全件
	r.GET("/sensors", func(c *gin.Context){

		//Execute query on TimescaleDB
		rows, err := dbpool.Query(ctx, "select * from sensor")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to execute query %v\n", err)
			return
		}
		defer rows.Close()
		fmt.Println("Successfully executed query", rows)
		
		// 返り値
		result := map[string]Sensor{}



		// for rows.Next() {
		// 	var sensor Sensor
		// 	err = rows.Scan(&sensor.sensor_id, &sensor.sensor_name, &sensor.user_id, &sensor.is_active, &sensor.created_at)
		// 	if err != nil {
		// 		fmt.Fprintf(os.Stderr, "Unable to scan %v\n", err)
		// 		return
		// 	}
		// 	fmt.Println(sensor.sensor_id, sensor.sensor_name, sensor.user_id, sensor.is_active, sensor.created_at)
		// 	result[sensor.sensor_id] = sensor
		// 	fmt.Println(result)
		// }
		// fmt.Println(result["1612e84d-734c-4b6c-93b7-d8c23c5099a8"])

		// 成功ログ
		c.JSON(http.StatusOK, result)
	})

	return r
}

func main() {

	// db connect
	ctx := context.Background()
	connStr := os.Getenv("CONNECT")
	dbpool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		panic(err)
	}
	defer dbpool.Close()

	r := setRouter(ctx, dbpool)
   
	r.Run(":8080")

}