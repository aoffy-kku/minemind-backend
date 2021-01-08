package service

import (
	"context"
	"fmt"
	"github.com/aoffy-kku/minemind-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

type AnalysisService struct {
	db  *mongo.Database
	col *mongo.Collection
}

func (a *AnalysisService) DeleteAnalysis(id primitive.ObjectID, uid string) error {
	ctx := context.Background()
	_, err := a.col.DeleteOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": id,
		},
		"user_id": bson.M{
			"$eq": uid,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *AnalysisService) GetAnalysisByDate(id string, date time.Time) ([]*model.AnalysisJSON, error) {
	ctx := context.Background()
	fromDate := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	toDate := fromDate.AddDate(0, 1, 0)
	opts := &options.FindOptions{}
	opts.SetSort(bson.M{
		"_id": -1,
	})
	cur, err := a.col.Find(ctx, bson.M{
		"user_id": bson.M{
			"$eq": id,
		},
		"created_at": bson.M{
			"$gte": fromDate,
			"$lt":  toDate,
		},
	}, opts)
	if err != nil {
		return nil, err
	}
	var docs []*model.AnalysisJSON
	for cur.Next(ctx) {
		var m model.Analysis
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		docs = append(docs, a.ToJSON(&m))
	}
	return docs, nil
}

func (a *AnalysisService) GetAnalysisByMode(id string, mode ...int64) ([]*model.AnalysisJSON, error) {
	ctx := context.Background()
	opts := &options.FindOptions{}
	opts.SetSort(bson.M{
		"_id": -1,
	})
	var docs []*model.AnalysisJSON
	var wg sync.WaitGroup
	for _, m := range mode {
		wg.Add(1)
		go func(md int64) {
			defer wg.Done()
			cur, err := a.col.Find(ctx, bson.M{
				"user_id": bson.M{
					"$eq": id,
				},
				"mode": bson.M{
					"$eq": md,
				},
			}, opts)
			if err != nil {
				return
			}
			for cur.Next(ctx) {
				var m model.Analysis
				if err := cur.Decode(&m); err != nil {
					return
				}
				docs = append(docs, a.ToJSON(&m))
			}
		}(m)
	}
	wg.Wait()
	return docs, nil
}

func (a *AnalysisService) CreateAnalysis(req model.CreateAnalysisRequestJSON) error {
	ctx := context.Background()
	now := time.Now()
	userCortisol := model.Cortisol{}
	userEvaluation := model.UserEvaluation{}
	mode := req.Mode
	opts := &options.FindOptions{}
	opts.SetSort(bson.M{
		"_id": -1,
	})
	opts.SetLimit(1)
	st5Value := int64(-1)
	phq9Value := int64(-1)
	cortisolValue := float64(-1)
	cortisolTimestamp := time.Now().Unix()
	// check st5
	if mode == 1 {
		cur, err := a.db.Collection("user_evaluation").Find(ctx, bson.M{
			"user_id": bson.M{
				"$eq": req.UserId,
			},
			"evaluation_id": bson.M{
				"$eq": "st5",
			},
		}, opts)
		if err != nil {
			return err
		}
		var result model.UserEvaluation
		for cur.Next(ctx) {
			if err := cur.Decode(&result); err != nil {
				return err
			}
			userEvaluation = result
		}
		st5Value = 0
		for _, question := range result.Questions {
			for _, answer := range question.Options {
				st5Value += answer.Value
			}
		}
	}
	// check phq9
	if mode == 2 {
		cur, err := a.db.Collection("user_evaluation").Find(ctx, bson.M{
			"user_id": bson.M{
				"$eq": req.UserId,
			},
			"evaluation_id": bson.M{
				"$eq": "phq9",
			},
		}, opts)
		if err != nil {
			return err
		}
		var result model.UserEvaluation
		for cur.Next(ctx) {
			if err := cur.Decode(&result); err != nil {
				return err
			}
			userEvaluation = result
		}
		phq9Value = 0
		for _, question := range result.Questions {
			for _, answer := range question.Options {
				phq9Value += answer.Value
			}
		}
	}
	// check cortisol
	if mode == 4 {
		cur, err := a.db.Collection("cortisol").Find(ctx, bson.M{
			"user_id": bson.M{
				"$eq": req.UserId,
			},
		}, opts)
		if err != nil {
			return err
		}
		for cur.Next(ctx) {
			var m model.Cortisol
			if err := cur.Decode(&m); err != nil {
				return err
			}
			userCortisol = m
			cortisolValue = m.Cortisol
		}
	}
	// check watch and cortisol
	if mode == 5 {
		cur, err := a.db.Collection("cortisol").Find(ctx, bson.M{
			"user_id": bson.M{
				"$eq": req.UserId,
			},
		}, opts)
		if err != nil {
			return err
		}
		for cur.Next(ctx) {
			var m model.Cortisol
			if err := cur.Decode(&m); err != nil {
				return err
			}
			userCortisol = m
			cortisolValue = m.Cortisol
			cortisolTimestamp = m.Timestamp
			break
		}
	}
	var user model.User
	if err := a.db.Collection("user").FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": req.UserId,
		},
	}).Decode(&user); err != nil {
		return err
	}
	res, err := a.col.InsertOne(ctx, &model.Analysis{
		Id:               primitive.NewObjectIDFromTimestamp(now),
		Mode:             req.Mode,
		UserId:           req.UserId,
		UserEvaluationId: userEvaluation.Id,
		UserEvaluation:   userEvaluation,
		CortisolId:       userCortisol.Id,
		Cortisol:         userCortisol,
		Status:           0,
		Class:            -1,
		Score:            -1,
		CreatedAt:        now,
		CreatedBy:        req.UserId,
		UpdatedAt:        now,
		UpdatedBy:        req.UserId,
	})
	if err != nil {
		return err
	}
	go func(id interface{}) {
		ct := time.Unix(cortisolTimestamp, 0)
		cmd := exec.Command(
			"python3",
			"./minemind_analysis/main.py",
			fmt.Sprintf("%d", mode),
			user.Email,
			user.WatchId,
			fmt.Sprintf("%d", st5Value),
			fmt.Sprintf("%d", phq9Value),
			fmt.Sprintf("%.2f", cortisolValue),
			fmt.Sprintf("\"%s\"", ct.Format("02/01/2006 03:04")),
			fmt.Sprintf("%s", user.BirthDate.Format("02/01/2006")),
			fmt.Sprintf("%s", user.Begin.Format("2006-01-02")),
			fmt.Sprintf("%s", time.Now().Format("2006-01-02")))
		log.Println(cmd)
		//stdout, err := cmd.StdoutPipe()
		//if err != nil {
		//	log.Println(err)
		//}
		//err = cmd.Start()
		//if err != nil {
		//	log.Println(err)
		//}
		//scanner := bufio.NewScanner(stdout)
		var result string
		//for scanner.Scan() {
		//	fmt.Println(scanner.Text())
		//	result = scanner.Text()
		//}
		//if err := cmd.Wait(); err != nil {
		//	log.Println(err)
		//}
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(err)
		}
		result = string(output)
		result = strings.TrimSpace(result)
		log.Println(result)
		data := strings.Split(result, " ")
		if len(data) != 4 {
			a.col.UpdateOne(ctx, bson.M{
				"_id": bson.M{
					"$eq": id,
				},
			}, bson.M{
				"$set": bson.M{
					"status": 2,
				},
			})
		} else {
			class := data[1]
			score := data[3]
			c, _ := strconv.ParseInt(class, 10, 64)
			s, _ := strconv.ParseFloat(score, 64)

			a.col.UpdateOne(ctx, bson.M{
				"_id": bson.M{
					"$eq": id,
				},
			}, bson.M{
				"$set": bson.M{
					"status": 1,
					"class":  c,
					"score":  s,
				},
			})
		}
	}(res.InsertedID)
	return nil
}

func (a *AnalysisService) ToJSON(m *model.Analysis) *model.AnalysisJSON {
	var questions []model.QuestionJSON
	for _, q := range m.UserEvaluation.Questions {
		var opts []model.OptionJSON
		for _, o := range q.Options {
			opts = append(opts, model.OptionJSON{
				Id:    o.Id,
				Title: o.Title,
				Value: o.Value,
			})
		}
		questions = append(questions, model.QuestionJSON{
			Id:      q.Id,
			Title:   q.Title,
			Options: opts,
		})
	}
	return &model.AnalysisJSON{
		Id:               m.Id,
		Mode:             m.Mode,
		UserId:           m.UserId,
		UserEvaluationId: m.UserEvaluationId,
		UserEvaluation: model.UserEvaluationJSON{
			Id:           m.UserEvaluation.Id,
			EvaluationId: m.UserEvaluation.EvaluationId,
			UserId:       m.UserEvaluation.UserId,
			Name:         m.UserEvaluation.Name,
			Description:  m.UserEvaluation.Description,
			Questions:    questions,
			CreatedAt:    m.CreatedAt,
			CreatedBy:    m.CreatedBy,
		},
		CortisolId: m.CortisolId,
		Cortisol: model.CortisolJSON{
			Id:        m.Cortisol.Id,
			UserId:    m.Cortisol.UserId,
			Cortisol:  m.Cortisol.Cortisol,
			Timestamp: m.Cortisol.Timestamp,
			CreatedAt: m.Cortisol.CreatedAt,
			CreatedBy: m.Cortisol.CreatedBy,
		},
		Status:    m.Status,
		Class:     m.Class,
		Score:     m.Score,
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedAt: m.UpdatedAt,
		UpdatedBy: m.UpdatedBy,
	}
}

func NewAnalysisService(db *mongo.Database) *AnalysisService {
	return &AnalysisService{
		db:  db,
		col: db.Collection("user_analysis"),
	}
}
