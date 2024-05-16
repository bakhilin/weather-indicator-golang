# Weather Forecast Application

This is a simple Golang application to fetch and display weather forecast data for a specific city using the OpenWeatherMap API.

## Prerequisites
Before running this application, you need to sign up on the OpenWeatherMap website (https://openweathermap.org/) to obtain an API key.

## Installation
1. Clone this repository or download the source code.
2. Replace "YOUR_API_KEY" in the getWeatherForecast function with your actual OpenWeatherMap API key.

## Usage
To run the application, execute the following command in your terminal:

bash
go run main.go


The application will fetch and display the weather forecast for the specified city (in this case, Moscow) for the next 5 days.

## Example Output

Date: 2021-10-01 12:00:00, Temperature: 289.45K, Description: clear sky
Date: 2021-10-02 12:00:00, Temperature: 291.37K, Description: few clouds
Date: 2021-10-03 12:00:00, Temperature: 290.65K, Description: light rain
Date: 2021-10-04 12:00:00, Temperature: 288.89K, Description: moderate rain
Date: 2021-10-05 12:00:00, Temperature: 285.14K, Description: light rain


The output includes the date and time of the forecast, temperature in Kelvin, and a brief description of the weather conditions.

Feel free to modify the city and days variables in the main function to fetch the weather forecast for a different city or a different number of days.
