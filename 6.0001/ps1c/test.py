def savings_calculator(annual_salary, savings_portion):
    monthly_salary = annual_salary/12
    semi_annual_raise = .07
    intrest_rate = .04
    current_savings = 0 
    month_count = 0 
    while month_count < 36 :
        month_count += 1
        current_savings += (monthly_salary*savings_portion) +(current_savings*intrest_rate/12)
        if month_count %6 == 0 :
            monthly_salary += monthly_salary *semi_annual_raise
    return current_savings


print(savings_calculator(150000, .441))