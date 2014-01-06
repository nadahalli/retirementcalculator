current_corpus = 0
annual_savings = 2
retirement_expenditure = 1
investment_returns = 12
inflation = 10
current_age = 34
death_age = 90

def savings(earning_years, current_corpus, annual_savings, investment_returns, inflation):
    total_savings = current_corpus
    for year in range(earning_years):
        total_savings += annual_savings
        total_savings *= (1 + investment_returns / 100.0)
        annual_savings *= (1 + inflation / 100.0)
    return total_savings

def spending(savings, spending_years, retirement_expenditure, investment_returns, inflation):
    for year in range(spending_years):
        savings -= retirement_expenditure
        retirement_expenditure *= (1 + inflation / 100.0)
        savings *= (1 + investment_returns / 100.0)
    return savings

def get_leftover_savings(age, earning_years, spending_years, current_corpus, annual_savings, retirement_expenditure, investment_returns, inflation):
    total_savings = savings(earning_years, current_corpus, annual_savings, investment_returns, inflation)
    leftover_savings = spending(total_savings, spending_years, retirement_expenditure, investment_returns, inflation)
    return total_savings, leftover_savings

if __name__ == '__main__':
    success = False
    for age in range(current_age, death_age):
        earning_years = age - current_age
        spending_years = death_age - age
        total_savings, leftover_savings = get_leftover_savings(age, 
                                                               earning_years, 
                                                               spending_years, 
                                                               current_corpus,
                                                               annual_savings, 
                                                               retirement_expenditure, 
                                                               investment_returns, 
                                                               inflation)
        if leftover_savings > 0:
            print age, retirement_expenditure, total_savings, leftover_savings
            success = True
            break
        print age, retirement_expenditure, total_savings, leftover_savings
        retirement_expenditure *= (1 + inflation / 100.0)
    if not success:
        print ':-('

