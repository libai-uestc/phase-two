package xorm

import (
	"fmt"
	"log/slog"

	"xorm.io/xorm"
)

func Read(engine *xorm.Engine) {
	user := User{City: "上海"}
	ok, err := engine.Select("*").Get(&user)
	if err != nil {
		slog.Error("读DB失败", "error", err)
	} else {
		if !ok {
			slog.Info("查无结果")
		} else {
			slog.Info("读库成功", "user", user)
		}
	}

	session := engine.Prepare()
	session.Select("*").Get(&user)

	var user2 User
	ok, err = engine.Cols("uid", "city", "gender", "keywords").IndexHint("force", "", "idx_uid").Where("uid>100 and degree='大专'").Where("degree='大专'").In("city", []string{"北京", "上海"}).And("degree like ?", "%科").Or("gender=?", "女").OrderBy("id desc").Desc("uid").Asc("city").Limit(1, 3).Get(&user2)
	if err != nil {
		slog.Error("读DB失败", "error", err)
	} else {
		if !ok {
			slog.Info("查无结果")
		} else {
			slog.Info("读库成功", "user", user2)
		}
	}

	var user3 *User
	_, err = engine.Get(user3)
	if err != nil {
		slog.Error("读DB失败", "error", err)
	}

	var user4 = new(User)
	_, err = engine.Get(user4)
	if err != nil {
		slog.Error("读DB失败", "error", err)
	} else {
		slog.Info("读库成功", "user", user4)
	}

	var users []User
	err = engine.Limit(3).Find(&users)
	if err != nil {
		slog.Error("读DB失败", "error", err)
	} else {
		fmt.Println("多个read结果")
		for _, u := range users {
			fmt.Printf("%+v\n", u)
		}
	}

	rows, err := engine.Limit(3).Rows(&User{})
	if err == nil {
		defer rows.Close()
		var u User
		fmt.Println("多个rows结果")
		for rows.Next() {
			rows.Scan(&u)
			fmt.Printf("user %+v\n", u)
		}
	} else {
		slog.Error("Rows failed", "error", err)
	}
}

func ReadWithStatistics(engine *xorm.Engine) {
	type Result struct {
		Degree string
		City   string
		Mid    float64
	}
	var results []Result

	err := engine.Table(User{}).Select("city, avg(id) as mid").GroupBy("city").Having("mid>0").Find(&results)
	if err == nil {
		fmt.Println("city mid")
		for _, result := range results {
			fmt.Printf("%s %.2f\n", result.City, result.Mid)
		}
	}
	err = engine.Distinct("degree,city").Table(User{}).Find(&results)
	if err == nil {
		fmt.Println("degree city")
		for _, result := range results {
			fmt.Printf("%s  %s\n", result.Degree, result.City)
		}
	}
	count, err := engine.Where("degree=?", "本科").Count(User{})
	if err == nil {
		fmt.Println("总数", count)
	}
}
