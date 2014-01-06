package hello

import (
    "fmt"
    "html/template"
    "net/http"
	"io/ioutil"
	"strconv"
	"os"
)

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/calculate", calculate)
}

func root(w http.ResponseWriter, r *http.Request) {
	form, err := ioutil.ReadFile("hello/form.html")
	if err != nil {
		panic (err)
	}
    fmt.Fprint(w, string(form))
}

type Result struct{
	RetirementAge string
}

func savings(earning_years float64, 
	current_corpus float64, 
	annual_savings float64, 
	investment_returns float64, 
	inflation float64) (float64) {

	total_savings := current_corpus
    for year := 0; year < int(earning_years); year++ {
		total_savings += annual_savings
		total_savings *= (1 + investment_returns / 100.0)
		annual_savings *= (1 + inflation / 100.0)
	}
    return total_savings
}

func spending(savings float64, 
	spending_years float64, 
	retirement_expenditure float64, 
	investment_returns float64, 
	inflation float64) (float64){

	for year := 0; year < int(spending_years); year++ {
		savings -= retirement_expenditure
		retirement_expenditure *= (1 + inflation / 100.0)
        savings *= (1 + investment_returns / 100.0)
	}
	return savings
}

func get_leftover_savings(age float64, 
	earning_years float64, 
	spending_years float64, 
	current_corpus float64, 
	annual_savings float64, 
	retirement_expenditure float64, 
	investment_returns float64, 
	inflation float64) (float64) {

	total_savings := savings(earning_years, current_corpus, annual_savings, investment_returns, inflation)
    leftover_savings := spending(total_savings, spending_years, retirement_expenditure, investment_returns, inflation)
    return leftover_savings
}

func retirementAgeCalculate(r *http.Request) (Result){
	current_corpus, _ := strconv.ParseFloat(r.FormValue("current_corpus"), 32)
	annual_savings, _ := strconv.ParseFloat(r.FormValue("annual_savings"), 32)
	retirement_expenditure, _ := strconv.ParseFloat(r.FormValue("retirement_expenditure"), 32)
	investment_returns, _ := strconv.ParseFloat(r.FormValue("investment_returns"), 32)
	inflation, _ := strconv.ParseFloat(r.FormValue("inflation"), 32)
	current_age, _ := strconv.ParseFloat(r.FormValue("current_age"), 32)
	death_age, _ := strconv.ParseFloat(r.FormValue("death_age"), 32)	
	
	for age := current_age; age < death_age; age += 1 {
		fmt.Fprint(os.Stdout, age)
		earning_years := age - current_age
		spending_years := death_age - age
		leftover_savings := get_leftover_savings(age, earning_years, spending_years, current_corpus, annual_savings, retirement_expenditure, investment_returns, inflation)
		if leftover_savings > 0 {
			return Result{RetirementAge: strconv.FormatFloat(age, 'f', 6, 32)}
		}
		retirement_expenditure *= (1 + inflation / 100.0)
	}
	return Result{RetirementAge: strconv.FormatFloat(death_age, 'f', 6, 32)}
}


func calculate(w http.ResponseWriter, r *http.Request) {
	result := retirementAgeCalculate(r)

	resultTemplateString, err := ioutil.ReadFile("hello/result.template")
	if err != nil {
		panic (err)
	}
	resultTemplate := template.New("calculate")
	resultTemplate, _ = resultTemplate.Parse(string(resultTemplateString))
	err = resultTemplate.Execute(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

