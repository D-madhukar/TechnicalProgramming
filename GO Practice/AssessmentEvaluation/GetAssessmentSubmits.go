package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"github.com/gorilla/mux"
	"log"
	"bytes"
	"io/ioutil"
)

type AssessmentSubmit struct {
	SchoolId	   string					     `json:"schoolid,omitempty"`
	// ID             string       				 `json:"_id"`
	AssessmentPost string      					 `json:"assessmentPost"`
	User           string        				 `json:"user"`
	Assessment     string       				 `json:"assessment"`
	Type           string       				 `json:"type"`
	Status         string       				 `json:"status"`
	TimeSpent      int         				     `json:"timeSpent"`
	Answer         string       				 `json:"answer,omitempty"`
	Post           string       				 `json:"post"`
	CreatedBy      string        				 `json:"created_by"`
	Options        []struct{
							Seqno    string `json:"seqno"`
							Text     string `json:"text"`
							Answer   bool   `json:"answer"`
							Isanswer bool   `json:"isanswer"`
					} 							 `json:"options,omitempty"`}
type Assessment struct {
	ID        string        		`json:"_id"`
	CreatedAt time.Time     		`json:"created_at"`
	UpdatedAt time.Time     		`json:"updated_at"`
	Seqno     int           		`json:"seqno"`
	Type      string        		`json:"type"`
	Answer    string        		`json:"answer,omitempty"`
	Marks     int           		`json:"marks"`
	Question  string        		`json:"question"`
	CreatedBy string        		`json:"created_by"`
	V         int           		`json:"__v"`
	Post      string        		`json:"post"`
	Options   []struct{
					Text   string `json:"text"`
					Seqno  string `json:"seqno"`
					Answer bool   `json:"answer"`
					ID     string `json:"_id"`
				}				 	`json:"options,omitempty"`
}
type SubmissionPost struct{
	Submittable		bool 		`json:"submittable"`
	Parent 			string 		`json:"parent"`
	Type			string 		`json:"type"`
	ResponseType	string		`json:"responseType"`
	School			string		`json:"school"`
	Class			string		`json:"class"`
	Author			string		`json:"author"`
	// Users			[]interface{}`json:"users"`
	// Poll            interface{}	`json:"poll"`
	Assessment		struct{
				Answered	int	 `json:"answered"`
				Correct		int	 `json:"correct"`
				Status		string`json:"status"`
				Timespent	int	 `json:"timespent"`
				Attempts	int  `json:"attempts"`
				Score 		int  `json:"score"`
			}					`json:"assessment"`
	Task					struct{
				Grade		int		 `json:"grade,omitempty"`
				Graded  bool	 `json:"graded,omitempty"`
		}						`json:"task,omitempty"`
	Stars		    int			`json:"stars"`
}

type SubmitId struct{
	Id string	`json:"id"`
}
var assessments []Assessment
var assessmentsubmits []AssessmentSubmit
var id SubmitId
var schoolid string
func main(){
	router:=mux.NewRouter()
	router.HandleFunc("/postassessmentsubmissions",getAssessmentSubmissions).Methods("POST")
	log.Fatal(http.ListenAndServe(":3132",router))
}
func getAssessmentSubmissions(w http.ResponseWriter,r *http.Request){
	json.NewDecoder(r.Body).Decode(&assessmentsubmits)
	schoolid=assessmentsubmits[0].SchoolId
	//Get the corresponding Assessments
	postid:=assessmentsubmits[0].AssessmentPost
	response,_:=http.Get("http://192.168.0.8:8080/api/school/"+schoolid+"getassessments/"+postid)
	body,_:=ioutil.ReadAll(response.Body)
	err1:=json.Unmarshal(body,&assessments)
	if err1!=nil{

	}else{

	}
	//Evaluate ,create and save the AssessmentPost
	var flag string="0"
	var submissionpostid string="123"
	var attempts int=0
	if(assessmentsubmits[0].Post!=""){
		flag="1"
		submissionpostid=assessmentsubmits[0].Post
		attempts=getAttempts(assessmentsubmits[0].User,assessmentsubmits[0].AssessmentPost)
	}
	notanswered,correct,score:=evaluateSubmissions()
	submissionpost:=prepareSubmissionPost(notanswered,correct,score,attempts)
	submissionpostjson,_:=json.Marshal(submissionpost)
	var url string="http://192.168.0.8:8080/api/school/"+schoolid+"savesubmissionpost/"+flag+"/"+submissionpostid
	fmt.Println(url)
	res,err:=http.Post(url,"application/json",bytes.NewBuffer(submissionpostjson))
	if err!=nil{
		fmt.Println("error",err);
	}else{
		fmt.Println("submissionpost sent to Node");
	}

	//Save the Assessment Submissions
	data,_:=ioutil.ReadAll(res.Body)
	err=json.Unmarshal(data,&id)
	initAssessmentSubmits(id.Id)
	jsondata,_:=json.Marshal(assessmentsubmits)
	_,err=http.Post("http://192.168.0.8:8080/api/school/"+schoolid+"savesubmissions","application/json",bytes.NewBuffer(jsondata))
	if err!=nil{
		fmt.Println("error in saving submissions  ",err);
	}else{
		fmt.Println("submissions sent successfully")
	}

}

//Getting previous attempt count
func getAttempts(userid string,postid string) int{
	var temppost SubmissionPost
	var url="http://192.168.0.8:8080/api/school/"+schoolid+"getassessmentpost/"+userid+"/"+postid
	res,_:=http.Get(url)
	body,_:=ioutil.ReadAll(res.Body)
	_=json.Unmarshal(body,&temppost)
	return temppost.Assessment.Attempts
}

//Assign the AssessmentPost id to Every corresponding submission
func initAssessmentSubmits(submitpostid string){
	for i,_:=range assessmentsubmits{
		assessmentsubmits[i].Post=submitpostid;//**
	}
}

//Creating the AssessmentPost Based on the Evaluation
func prepareSubmissionPost(notanswered,correct,score,attempts int) SubmissionPost{
	var submissionpost SubmissionPost
	answered:=len(assessments)-notanswered
	submissionpost.Submittable=true
	submissionpost.Parent=assessmentsubmits[0].AssessmentPost//**
	submissionpost.Type="Assessment"
	submissionpost.ResponseType="submission"
	submissionpost.Author=assessmentsubmits[0].User
	submissionpost.Assessment.Answered=answered
	submissionpost.Assessment.Correct=correct
	submissionpost.Assessment.Status="success"
	submissionpost.Assessment.Timespent=30
	submissionpost.Assessment.Attempts=attempts+1
	submissionpost.Assessment.Score=score
	submissionpost.Task.Grade=score
	submissionpost.Stars=1
	return submissionpost
}

//Evaluating the AssessmentSubmissions
func evaluateSubmissions()(int,int,int){
	var notanswered int=0
	var correct int =0
	var score int=0
	for i,submission:=range assessmentsubmits{
			switch submission.Type{
				case "TF":{
					if submission.Answer=="not_answered"{
						notanswered++
						assessmentsubmits[i].Status="not answered"
					}else if submission.Answer==assessments[i].Answer{
						correct++
						score=score+assessments[i].Marks
						assessmentsubmits[i].Status="correct"
					}else{
						assessmentsubmits[i].Status="not correct"
					}
				}
				case "SA":{
					if submission.Answer=="not_answered"{
						notanswered++
						assessmentsubmits[i].Status="not answered"
					}else if checkShortAnswer(i){
						correct++
						score=score+assessments[i].Marks
						assessmentsubmits[i].Status="correct"
					}else{
						assessmentsubmits[i].Status="not correct"
					}
				}
				case "MC":{
					if checkMcAttempt(i){
						notanswered++
						assessmentsubmits[i].Status="not answered"
					}else if checkMultipleChoice(i){
						correct++
						score=score+assessments[i].Marks
						assessmentsubmits[i].Status="correct"
					}else{
						assessmentsubmits[i].Status="not correct"
					}
				}
			}

	}
	return notanswered,correct,score
}

//Checking whether Student attempted the MC Question or not.
func checkMcAttempt(i int) bool{
	for _,option:=range assessmentsubmits[i].Options{
		if option.Answer==true{
			return false
		}
	}
	return true
}

//Checking SA correctness
func checkShortAnswer(i int ) bool{
	var studentanswer=assessmentsubmits[i].Answer
	for _,option:=range assessments[i].Options{
		if option.Text==studentanswer{
			return true
		}
	}
	return false
}

//checking MC correctness
func checkMultipleChoice(i int)bool{
	for j,option:=range assessmentsubmits[i].Options{
		if option.Answer!=assessments[i].Options[j].Answer{
			return false
		}
	}
	return true
}
