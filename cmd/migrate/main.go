package main

import (
	"context"
	"fmt"
	db2 "github.com/aoffy-kku/minemind-backend/db"
	"github.com/aoffy-kku/minemind-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var db = db2.New()
var ctx = context.Background()

func main()  {
	//migrateRole()
	//migrateST5()
	//migratePHQ9()
	//migrateMood()
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