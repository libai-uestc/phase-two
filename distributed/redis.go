package distributed

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func StringValue(ctx context.Context, client *redis.Client) {
	key := "name"
	value := "libai"
	defer client.Del(ctx, key)

	err := client.Set(ctx, key, value, 1*time.Second).Err()
	checkError(err)

	client.Expire(ctx, key, 3*time.Second)
	time.Sleep(2 * time.Second)

	v2, err := client.Get(ctx, key).Result()
	checkError(err)
	fmt.Println(v2)

	err = client.Set(ctx, "age", 18, 1*time.Second).Err()
	checkError(err)
	v3, err := client.Get(ctx, "age").Int()
	checkError(err)
	fmt.Printf("age=%d\n", v3)
}

type Student struct {
	Id   int
	Name string
}

func WriteStudent2Redis(client *redis.Client, stu *Student) error {
	if stu == nil {
		return nil
	}
	ctx := context.Background()
	key := "STU_" + strconv.Itoa(stu.Id)
	v, err := json.Marshal(stu)
	if err != nil {
		return err
	}
	err = client.Set(ctx, key, string(v), 5*time.Minute).Err()
	return err
}

func GetStudentFromRedis(client *redis.Client, sid int) *Student {
	ctx := context.Background()
	key := "STU_" + strconv.Itoa(sid)
	v, err := client.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Println(err)
		}
		return nil
	}
	var stu Student
	err = json.Unmarshal([]byte(v), &stu)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &stu
}

func DeleteKey(ctx context.Context, client *redis.Client) {
	n, err := client.Del(ctx, "not_exists").Result()
	if err == nil {
		fmt.Printf("删除%d个key\n", n)
	}
}

func ListValue(ctx context.Context, client *redis.Client) {
	key := "ids"
	defer client.Del(ctx, key)

	values := []interface{}{1, "李", 3, 4, 3, 1}
	err := client.RPush(ctx, key, values...).Err()
	checkError(err)

	v2, err := client.LRange(ctx, key, 0, -1).Result()
	checkError(err)
	fmt.Println(v2)
}

func SetValue(ctx context.Context, client *redis.Client) {
	key := "ids"
	defer client.Del(ctx, key)

	values := []interface{}{1, "李", 3, 4, 3, 1}
	err := client.SAdd(ctx, key, values...).Err()
	checkError(err)

	var value any
	value = 1
	if client.SIsMember(ctx, key, value).Val() {
		fmt.Printf("Set中包含%#v\n", value)
	} else {
		fmt.Printf("Set中不包含%#v\n", value)
	}
	value = "1"
	if client.SIsMember(ctx, key, value).Val() {
		fmt.Printf("Set中包含%#v\n", value)
	} else {
		fmt.Printf("Set中不包含%#v\n", value)
	}
	value = 2
	if client.SIsMember(ctx, key, value).Val() {
		fmt.Printf("Set中包含%#v\n", value)
	} else {
		fmt.Printf("Set中不包含%#v\n", value)
	}
	for _, ele := range client.SMembers(ctx, key).Val() {
		fmt.Println(ele)
	}

	key2 := "ids2"
	defer client.Del(ctx, key2)
	values = []interface{}{1, "李", "白", 18}
	err = client.SAdd(ctx, key2, values...).Err()
	checkError(err)

	fmt.Println("key1 - key2 差集")
	for _, ele := range client.SDiff(ctx, key, key2).Val() {
		fmt.Println(ele)
	}
	fmt.Println("key2 - key1 差集")
	for _, ele := range client.SDiff(ctx, key2, key).Val() {
		fmt.Println(ele)
	}

	fmt.Println("key & key2 交集")
	for _, ele := range client.SInter(ctx, key, key2).Val() {
		fmt.Println(ele)
	}
}

func ZsetValue(ctx context.Context, client *redis.Client) {
	key := "ids"
	defer client.Del(ctx, key)

	values := []redis.Z{{Member: "张三", Score: 77.7}, {Member: "李四", Score: 99.9}, {Member: "王五", Score: 88.8}}
	err := client.ZAdd(ctx, key, values...).Err()
	checkError(err)

	for _, ele := range client.ZRange(ctx, key, 0, -1).Val() {
		fmt.Println(ele)
	}
}

func HashtableValue(ctx context.Context, client *redis.Client) {
	student1 := map[string]interface{}{"Name": "张三", "Age": 18, "Height": 173.5}
	err := client.HMSet(ctx, "学生1", student1).Err()
	checkError(err)
	student2 := map[string]interface{}{"Name": "李四", "Age": 28, "Height": 199.9}
	err = client.HMSet(ctx, "学生2", student2).Err()
	checkError(err)

	age, err := client.HGet(ctx, "学生2", "Age").Int()
	checkError(err)
	fmt.Printf("age=%d\n", age)

	for field, value := range client.HGetAll(ctx, "学生1").Val() {
		fmt.Printf("field:%s  value:%s\n", field, value)
	}

	client.Del(ctx, "学生1")
	client.Del(ctx, "学生2")
}

func checkError(err error) {
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key不存在")
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func Scan(ctx context.Context, client *redis.Client) {
	if client == nil {
		log.Printf("connect redis failed")
		os.Exit(1)
	}
	const (
		MID = "_libai_"
	)
	for i := 0; i < 10; i++ {
		key := strconv.Itoa(i) + MID + strconv.Itoa(i)
		err := client.Set(ctx, key, "1", 0).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
	const COUNT = 100
	var cursor uint64 = 0
	dup := make(map[string]struct{}, 10)
	for {
		keys, c, err := client.Scan(ctx, cursor, "*"+MID+"*", COUNT).Result()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("cursor %d keys count %d\n", c, len(keys))
		for _, key := range keys {
			dup[key] = struct{}{}
		}
		if c == 0 {
			break
		}
		cursor = c
	}
	fmt.Println("total", len(dup))
	for key := range dup {
		fmt.Println(key)
	}
}
