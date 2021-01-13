package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	db2 "github.com/aoffy-kku/minemind-backend/db"
	"github.com/aoffy-kku/minemind-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var db = db2.New()
var ctx = context.Background()

func main()  {
	//migrateRole()
	//migrateST5()
	//migratePHQ9()
	//migrateMood()
	//migrateVersion()
	migrateTester()
}

func migrateTester() {
	ids := []string{
		"CLM_26qxn",
		"CLM_0SIM5",
		"CLM_CuXIM",
		"CLM_wA2Z7",
		"CLM_thtiL",
		"CLM_fyZ1L",
		"CLM_KJPeA",
		"CLM_5jkwP",
		"CLM_kbY4W",
		"CLM_mSs8Q",
	}
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTAzNDAwMzcsImlkIjoiYWRtaW5AbWluZW1pbmQubmV0In0.vEl1BOk7vXGT1UZ88O9GawtpQhu7r4Jb3vfuWrqnxkE"
	for i, id := range ids {
		prefix := ""
		if i + 1 < 10 {
			prefix = fmt.Sprintf("minemind00%d.kku", i + 1)
		} else {
			prefix = fmt.Sprintf("minemind0%d.kku", i + 1)
		}
		user := model.CreateUserRequestJSON{
			Email:      fmt.Sprintf("%s@minemind.net", prefix),
			Password:   prefix,
			DisplayName: prefix,
			WatchId:     id,
		}
		body, err := json.Marshal(&user)
		if err != nil {
			panic(err)
		}
		req, err := http.NewRequest("POST", "https://api.minemind.net/v1/users", bytes.NewBuffer(body))
		if err != nil {
			panic(err)
		}
		req.Header.Add("authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Add("accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		res, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(res))
	}
}

func migrateVersion() {
	type Version struct {
		Id string `bson:"_id" json:"id"`
		Code string `bson:"code" json:"code"`
		Url string  `bson:"url" json:"url"`
	}
	col := db.Collection("version")
	result, err := col.InsertOne(ctx, &Version{
		Id: "minemind",
		Code: "1.0.0",
		Url:  "",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted %s \n", result.InsertedID)
}


func migrateRole() {
	col := db.Collection("role")
	result, err := col.InsertMany(ctx, []interface{}{
		model.Role{
			Id:        "admin",
			CreatedAt: time.Now(),
			CreatedBy: "system",
		},
		model.Role{
			Id:        "officer",
			CreatedAt: time.Now(),
			CreatedBy: "system",
		},
		model.Role{
			Id:        "user",
			CreatedAt: time.Now(),
			CreatedBy: "system",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted %d rows\n", len(result.InsertedIDs))
}

func migrateST5() {
	col := db.Collection("evaluation")
	ts := time.Now()
	opts := []model.Option{
		model.Option{
			Id:         1,
			Title:      "เป็นน้อยมากหรือแทบไม่มี",
			Value:      0,
		},
		model.Option{
			Id:         2,
			Title:      "เป็นบางครั้ง",
			Value:      1,
		},
		model.Option{
			Id:         3,
			Title:      "เป็นบ่อยครั้ง",
			Value:      2,
		},
		model.Option{
			Id:         4,
			Title:      "เป็นประจำ",
			Value:      3,
		},
	}
	result, err := col.InsertOne(ctx, &model.Evaluation{
		Id:          "st5",
		Name:        "แบบประเมินความเครียด (ST- 5)",
		Description: "อาการหรือความรู้สึกที่เกิดในระยะ 2 - 4 สัปดาห์",
		Questions:   []model.Question{
			model.Question{
				Id:           1,
				Title:        "มีปัญหาการนอน นอนไม่หลับหรือนอนมาก",
				Options:      opts,
			},
			model.Question{
				Id:           2,
				Title:        "มีสมาธิน้อยลง",
				Options:      opts,
			},
			model.Question{
				Id:           3,
				Title:        "หงุดหงิด / กระวนกระวาย / ว้าวุ้นใจ",
				Options:      opts,
			},
			model.Question{
				Id:           4,
				Title:        "รู้สึกเบื่อ เซ็ง",
				Options:      opts,
			},
			model.Question{
				Id:           5,
				Title:        "ไม่อยากพบปะผู้คน",
				Options:      opts,
			},
		},
		CreatedAt:   ts,
		CreatedBy:   "system",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}

func migratePHQ9() {
	col := db.Collection("evaluation")
	ts := time.Now()
	opts := []model.Option{
		model.Option{
			Id:         1,
			Title:      "ไม่เลย",
			Value:      0,
		},
		model.Option{
			Id:         2,
			Title:      "มีบางวันหรือไม่บ่อย",
			Value:      1,
		},
		model.Option{
			Id:         3,
			Title:      "มีค่อนข้างบ่อย",
			Value:      2,
		},
		model.Option{
			Id:         4,
			Title:      "มีเกือบทุกวัน",
			Value:      3,
		},
	}
	result, err := col.InsertOne(ctx, &model.Evaluation{
		Id:          "phq9",
		Name:        "แบบทดสอบภาวะซึมเศร้า PHQ-9",
		Description: "ในช่วง 2 สัปดาห์ที่ผ่านมา ท่านมีอาการดังต่อไปนี้บ่อยแค่ไหน?",
		Questions:   []model.Question{
			model.Question{
				Id:           1,
				Title:        "เบื่อ ทำอะไร ๆ ก็ไม่เพลิดเพลิน",
				Options:      opts,
			},
			model.Question{
				Id:           2,
				Title:        "ไม่สบายใจ ซึมเศร้า หรือท้อแท้",
				Options:      opts,
			},
			model.Question{
				Id:           3,
				Title:        "หลับยาก หรือหลับ ๆ ตื่น ๆ หรือหลับมากไป",
				Options:      opts,
			},
			model.Question{
				Id:           4,
				Title:        "เหนื่อยง่าย หรือไม่ค่อยมีแรง",
				Options:      opts,
			},
			model.Question{
				Id:           5,
				Title:        "เบื่ออาหาร หรือกินมากเกินไป",
				Options:      opts,
			},
			model.Question{
				Id:           6,
				Title:        "รู้สึกไม่ดีกับตัวเอง คิดว่าตัวเองล้มเหลว หรือเป็นคนทำให้ตัวเอง หรือครอบครัวผิดหวัง",
				Options:      opts,
			},
			model.Question{
				Id:           7,
				Title:        "สมาธิไม่ดีเวลาทำอะไร เช่น ดูโทรทัศน์ ฟังวิทยุ หรือทำงานที่ต้องใช้ความตั้งใจ",
				Options:      opts,
			},
			model.Question{
				Id:           8,
				Title:        "พูดหรือทำอะไรช้าจนคนอื่นมองเห็น หรือกระสับกระส่ายจนท่านอยู่ไม่นิ่งเหมือนเคย",
				Options:      opts,
			},
			model.Question{
				Id:           9,
				Title:        "คิดทำร้ายตนเอง หรือคิดว่าถ้าตาย ๆ ไปเสียคงจะดี",
				Options:      opts,
			},
		},
		CreatedAt:   ts,
		CreatedBy:   "system",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}

func migrateMood() {
	moods := []string{
		"เซ็ง",
		"เบื่อ",
		"เศร้า",
		"เหนื่อย",
		"หมดหวัง",
		"โกรธ",
		"เกลียด",
		"หงุดหงิด",
		"ไร้ค่า",
		"เครียด",
		"เหงา",
		"พอกันที",
		"ไม่สน",
		"โดดเดี่ยว",
		"ง่วง",
		"ปวดหัว",
		"กังวล",
		"ตื่นเต้น",
		"รำคาญ",
		"สุข",
		"กลัว",
		"ว่างเปล่า",
		"ดี",
		"ใจสั่น",
		"เหนื่อยใจ",
		"ท้อแท้",
		"อ่อนเพลีย",
		"สบายใจ",
	}
	var docs []interface{}
	for _, mood := range moods {
		ts := time.Now()
		docs = append(docs, model.Mood{
			Id:        primitive.NewObjectIDFromTimestamp(ts),
			Name:      mood,
			CreatedAt: ts,
			CreatedBy: "system",
		})
	}
	col := db.Collection("mood")
	result, err := col.InsertMany(ctx, docs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted %d rows\n", len(result.InsertedIDs))
}